package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/builder"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type onboardController struct {
	*baseController
	googleConf *config.Google
}

func NewOnboardController(googleConf *config.Google) (port.OnboardController, error) {
	if googleConf == nil {
		return nil, errorx.NilPointerErr{Item: "google config"}
	}
	return &onboardController{baseController: &baseController{}, googleConf: googleConf}, nil
}

func (oc onboardController) SignupWith(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")
	if provider == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var url bytes.Buffer

	switch provider {
	case "google":
		url.WriteString(oc.googleConf.OAuth2Url)

		queryBuilder := builder.NewQuery()
		queryBuilder.Add("client_id", oc.googleConf.ClientId)
		queryBuilder.Add("redirect_uri", oc.googleConf.RedirectCodeUrl)
		queryBuilder.Add("response_type", oc.googleConf.ResponseType)
		queryBuilder.Add("scope", strings.Join(oc.googleConf.Scopes, " "))
		queryBuilder.Add("access_type", oc.googleConf.AccessType)
		queryBuilder.Add(
			"include_granted_scopes", strconv.FormatBool(oc.googleConf.IncludeGrantedScopes),
		)

		url.WriteString(queryBuilder.Build())
		break
	default:
		err := payload.HTTPError{
			StatusCode:    http.StatusNotFound,
			StatusMessage: fmt.Sprintf("Provider %s not found", provider),
		}
		oc.handleError(w, err)
	}

	sso := payload.SSO{
		Provider: provider,
		Url:      url.String(),
	}
	err := oc.handleResponse(w, sso)
	if err != nil {
		logger.ERR.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (oc onboardController) LoginWith(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")
	if provider == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (oc onboardController) RedirectGoogle(w http.ResponseWriter, r *http.Request) {
	redirectErr := r.URL.Query().Get("error")
	if redirectErr != "" {
		err := payload.HTTPError{
			StatusCode:    http.StatusForbidden,
			StatusMessage: redirectErr,
		}
		oc.handleError(w, err)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		err := payload.HTTPError{
			StatusCode:    http.StatusNotFound,
			StatusMessage: "authentication code not found in query params",
		}
		oc.handleError(w, err)
		return
	}

	var url bytes.Buffer

	url.WriteString(oc.googleConf.TokenUrl)

	queryBuilder := builder.NewQuery()
	queryBuilder.Add("code", code)
	queryBuilder.Add("client_id", oc.googleConf.ClientId)
	queryBuilder.Add("client_secret", oc.googleConf.ClientSecret)
	queryBuilder.Add("redirect_uri", oc.googleConf.RedirectTokenUrl)
	queryBuilder.Add("grant_type", "authorization_code")

	url.WriteString(queryBuilder.Build())

	go func() {
		_, err := http.Post(url.String(), "application/x-www-form-urlencoded", nil)
		if err != nil {
			logger.ERR.Println(err)
		}
	}()
}

func (oc onboardController) RedirectGoogleToken(w http.ResponseWriter, r *http.Request) {
	var (
		tokenResp payload.GoogleTokenResp
		data      []byte
		err       error
	)

	data, err = io.ReadAll(r.Body)
	if err != nil {
		logger.ERR.Println(err)
		http.Redirect(w, r, config.Root.String(), http.StatusSeeOther)
		return
	}
	err = json.Unmarshal(data, &tokenResp)
	if err != nil {
		logger.ERR.Println(err)
		http.Redirect(w, r, config.Root.String(), http.StatusSeeOther)
		return
	}
	logger.INFO.Println(tokenResp)
}

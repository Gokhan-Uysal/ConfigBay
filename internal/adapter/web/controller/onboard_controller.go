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
		queryBuilder.Add("redirect_uri", oc.googleConf.RedirectUrl)
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
	var (
		tokenResp *payload.GoogleTokenResp
		code      string
		cookie    *http.Cookie
		err       error
	)
	errMessage := r.URL.Query().Get("error")
	if errMessage != "" {
		err := payload.HTTPError{
			StatusCode:    http.StatusForbidden,
			StatusMessage: errMessage,
		}
		oc.handleError(w, err)
		return
	}

	code = r.URL.Query().Get("code")
	if code == "" {
		err := payload.HTTPError{
			StatusCode:    http.StatusNotFound,
			StatusMessage: "authentication code not found in query params",
		}
		oc.handleError(w, err)
		return
	}

	tokenResp, err = oc.fetchToken(code)
	if err != nil {
		err := payload.HTTPError{
			StatusCode:    http.StatusNotFound,
			StatusMessage: err.Error(),
		}
		oc.handleError(w, err)
		return
	}
	cookie = &http.Cookie{
		Name:     "token",
		Value:    tokenResp.AccessToken,
		Path:     config.Home.String(),
		MaxAge:   tokenResp.ExpiresIn,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, cookie)
	r.AddCookie(cookie)

	http.Redirect(w, r, config.Home.String(), http.StatusSeeOther)
}

func (oc onboardController) fetchToken(code string) (*payload.GoogleTokenResp, error) {
	var (
		resp      *http.Response
		url       bytes.Buffer
		tokenResp payload.GoogleTokenResp
		data      []byte
		err       error
	)

	url.WriteString(oc.googleConf.TokenUrl)

	queryBuilder := builder.NewQuery()
	queryBuilder.Add("code", code)
	queryBuilder.Add("client_id", oc.googleConf.ClientId)
	queryBuilder.Add("client_secret", oc.googleConf.ClientSecret)
	queryBuilder.Add("redirect_uri", oc.googleConf.RedirectUrl)
	queryBuilder.Add("grant_type", "authorization_code")

	url.WriteString(queryBuilder.Build())
	resp, err = http.Post(url.String(), "application/x-www-form-urlencoded", nil)
	if err != nil {
		return nil, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tokenResp)
	if err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

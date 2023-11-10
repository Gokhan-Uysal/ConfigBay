package controller

import (
	"bytes"
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/builder"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
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
	var buffer bytes.Buffer

	switch provider {
	case "google":
		buffer.WriteString(oc.googleConf.OAuth2)

		queryBuilder := builder.NewQuery()
		queryBuilder.Add("client_id", oc.googleConf.ClientId)
		queryBuilder.Add("redirect_uri", oc.googleConf.RedirectUrl)
		queryBuilder.Add("response_type", oc.googleConf.ResponseType)
		queryBuilder.Add("scope", strings.Join(oc.googleConf.Scopes, " "))
		queryBuilder.Add("access_type", oc.googleConf.AccessType)
		queryBuilder.Add(
			"include_granted_scopes", strconv.FormatBool(oc.googleConf.IncludeGrantedScopes),
		)

		buffer.WriteString(queryBuilder.Build())
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
		Url:      buffer.String(),
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
			StatusMessage: "Authentication code not found in query params",
		}
		oc.handleError(w, err)
		return
	}

	ssoCookie := &http.Cookie{
		Name:     "CODE",
		Value:    code,
		Secure:   true,
		HttpOnly: true,
		MaxAge:   60 * 20,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, ssoCookie)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

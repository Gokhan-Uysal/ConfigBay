package controller

import (
	"fmt"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
	"net/http"
)

type onboardController struct {
	*baseController
	googleAuthService port.GoogleAuthService
}

func NewOnboardController(googleAuthService port.GoogleAuthService) (port.OnboardController, error) {
	if googleAuthService == nil {
		return nil, errorx.NilPointerErr{Item: "google authentication service"}
	}
	return &onboardController{
		baseController:    &baseController{},
		googleAuthService: googleAuthService,
	}, nil
}

func (oc onboardController) SignupWith(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")
	if provider == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var (
		sso payload.SSO
		err error
	)

	switch provider {
	case "google":
		sso = oc.googleAuthService.BuildSSO(provider)
		break
	default:
		err := payload.HTTPError{
			StatusCode:    http.StatusNotFound,
			StatusMessage: fmt.Sprintf("Provider %s not found", provider),
		}
		oc.handleError(w, err)
	}

	err = oc.handleResponse(w, sso)
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
		tokenResp *payload.GoogleToken
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

	tokenResp, err = oc.googleAuthService.FetchToken(code)
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

package controller

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"net/http"
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

	w = oc.enableCors(w)
	w = oc.addCors(w, oc.googleConf.OAuth2)
}

func (oc onboardController) LoginWith(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")
	if provider == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w = oc.enableCors(w)
	w = oc.addCors(w, oc.googleConf.OAuth2)
}

package controller

import (
	"bytes"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/builder"
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
	http.Redirect(w, r, buffer.String(), http.StatusFound)
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

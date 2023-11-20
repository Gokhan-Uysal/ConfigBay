package auth

import (
	"bytes"
	"encoding/json"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/port"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/builder"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type googleAuthService struct {
	googleConf *config.Google
}

func NewGoogleAuthService(googleConf *config.Google) (port.GoogleAuthService, error) {
	if googleConf == nil {
		return nil, errorx.NilPointerErr{Item: "google config"}
	}
	return &googleAuthService{googleConf: googleConf}, nil
}

func (gs googleAuthService) BuildSSO(provider string) payload.SSO {
	var (
		url bytes.Buffer
	)

	url.WriteString(gs.googleConf.OAuth2Url)

	queryBuilder := builder.NewQuery()
	queryBuilder.Add("client_id", gs.googleConf.ClientId)
	queryBuilder.Add("redirect_uri", gs.googleConf.RedirectUrl)
	queryBuilder.Add("response_type", gs.googleConf.ResponseType)
	queryBuilder.Add("scope", strings.Join(gs.googleConf.Scopes, " "))
	queryBuilder.Add("access_type", gs.googleConf.AccessType)
	queryBuilder.Add(
		"include_granted_scopes", strconv.FormatBool(gs.googleConf.IncludeGrantedScopes),
	)

	url.WriteString(queryBuilder.Build())

	return payload.SSO{
		Provider: provider,
		Url:      url.String(),
	}
}

func (gs googleAuthService) FetchToken(code string) (*payload.GoogleToken, error) {
	var (
		resp      *http.Response
		url       bytes.Buffer
		tokenResp payload.GoogleToken
		data      []byte
		err       error
	)

	url.WriteString(gs.googleConf.ApiUrl)
	url.WriteString("/token")

	queryBuilder := builder.NewQuery()
	queryBuilder.Add("code", code)
	queryBuilder.Add("client_id", gs.googleConf.ClientId)
	queryBuilder.Add("client_secret", gs.googleConf.ClientSecret)
	queryBuilder.Add("redirect_uri", gs.googleConf.RedirectUrl)
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

func (gs googleAuthService) RefreshToken(refreshToken string) (*payload.GoogleToken, error) {
	var (
		resp      *http.Response
		url       bytes.Buffer
		tokenResp payload.GoogleToken
		data      []byte
		err       error
	)

	url.WriteString(gs.googleConf.ApiUrl)
	url.WriteString("/token")

	queryBuilder := builder.NewQuery()
	queryBuilder.Add("client_id", gs.googleConf.ClientId)
	queryBuilder.Add("client_secret", gs.googleConf.ClientSecret)
	queryBuilder.Add("refresh_token", refreshToken)
	queryBuilder.Add("grant_type", "refresh_token")

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

func (gs googleAuthService) RevokeToken(token string) error {
	var (
		url bytes.Buffer
		err error
	)

	url.WriteString(gs.googleConf.ApiUrl)
	url.WriteString("/revoke")

	queryBuilder := builder.NewQuery()
	queryBuilder.Add("token", token)

	url.WriteString(queryBuilder.Build())
	_, err = http.Post(url.String(), "application/x-www-form-urlencoded", nil)

	return err
}

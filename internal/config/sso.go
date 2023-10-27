package config

type Google struct {
	ClientId     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	OAuth2       string   `json:"oauth2"`
	ResponseType string   `json:"responseType"`
	RedirectUrl  string   `json:"redirectUrl"`
	Scopes       []string `json:"scopes"`
}

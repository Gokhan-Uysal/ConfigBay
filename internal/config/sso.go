package config

type SSOProvider struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Google struct {
	ClientId             string   `json:"clientId"`
	ClientSecret         string   `json:"clientSecret"`
	OAuth2Url            string   `json:"OAuth2Url"`
	ApiUrl               string   `json:"apiUrl"`
	ResponseType         string   `json:"responseType"`
	RedirectUrl          string   `json:"redirectUrl"`
	Scopes               []string `json:"scopes"`
	AccessType           string   `json:"accessType"`
	IncludeGrantedScopes bool     `json:"includeGrantedScopes"`
}

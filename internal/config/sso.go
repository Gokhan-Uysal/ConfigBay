package config

type SSOProvider struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Google struct {
	ClientId             string   `json:"clientId"`
	ClientSecret         string   `json:"clientSecret"`
	OAuth2Url            string   `json:"OAuth2Url"`
	TokenUrl             string   `json:"tokenUrl"`
	ResponseType         string   `json:"responseType"`
	RedirectCodeUrl      string   `json:"redirectCodeUrl"`
	RedirectTokenUrl     string   `json:"redirectTokenUrl"`
	Scopes               []string `json:"scopes"`
	AccessType           string   `json:"accessType"`
	IncludeGrantedScopes bool     `json:"includeGrantedScopes"`
}

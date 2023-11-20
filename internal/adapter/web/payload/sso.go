package payload

type SSO struct {
	Provider string `json:"provider"`
	Url      string `json:"url"`
}

type GoogleTokenResp struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenId      string `json:"id_token"`
}

package payload

type SSO struct {
	Provider string `json:"provider"`
	Url      string `json:"url"`
}

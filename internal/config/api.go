package config

type Api struct {
	Version      string        `json:"version"`
	Name         string        `json:"name"`
	Host         string        `json:"host"`
	Port         int           `json:"port"`
	LogLevel     string        `json:"logLevel"`
	Template     string        `json:"template"`
	Static       string        `json:"static"`
	SSOProviders []SSOProvider `json:"ssoProviders"`
}

type Endpoint string

func (e Endpoint) String() string {
	return string(e)
}

const (
	Root                Endpoint = "/"
	Home                Endpoint = "/home"
	Signup              Endpoint = "/signup"
	Login               Endpoint = "/login"
	SignupWith          Endpoint = "/signup-with"
	LoginWith           Endpoint = "/login-with"
	RedirectGoogle      Endpoint = "/redirect/google"
	RedirectGoogleToken Endpoint = "/redirect/google/token"
)

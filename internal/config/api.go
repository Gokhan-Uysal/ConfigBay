package config

type Api struct {
	Version  string `json:"version"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	LogLevel string `json:"logLevel"`
	Template string `json:"template"`
	Static   string `json:"static"`
}

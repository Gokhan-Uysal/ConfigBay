package config

type Api struct {
	Version  string `json:"version"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Template string `json:"template"`
	Static   string `json:"static"`
}

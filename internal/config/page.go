package config

type page struct {
	Name         string   `json:"name"`
	HasNavbar    bool     `json:"hasNavbar"`
	HasFooter    bool     `json:"hasFooter"`
	AuthRequired bool     `json:"authRequired"`
	Navbar       []Navbar `json:"navbar"`
}

type RootPage struct {
	page
}

type OnboardPage struct {
	page
	Providers []Provider `json:"providers"`
}

type HomePage struct {
	page
}

type Navbar struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

type Provider struct {
	Provider string `json:"provider"`
	Icon     string `json:"icon"`
}

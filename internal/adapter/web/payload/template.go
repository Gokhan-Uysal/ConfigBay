package payload

import "time"

type NavbarItem struct {
	Href  string
	Label string
}

type ProjectItem struct {
	Icon        []byte
	Title       string
	Description string
	GroupCount  int
	UserCount   int
	CreatedAt   time.Time
}

type OnboardItem struct {
	Access   string
	Provider string
	Icon     string
}

type RootPage struct {
	NavbarItems []NavbarItem
}

type HomePage struct {
	NavbarItems  []NavbarItem
	ProjectItems []ProjectItem
}

type OnboardPage struct {
	NavbarItems  []NavbarItem
	Access       string
	OnboardItems []OnboardItem
}

var homeItem = NavbarItem{Href: "/home", Label: "Home"}
var projectItem = NavbarItem{Href: "/projects", Label: "Projects"}
var HomePageNavbar = []NavbarItem{homeItem, projectItem}

var loginItem = NavbarItem{Href: "/login", Label: "Login"}
var signupItem = NavbarItem{Href: "/signup", Label: "Signup"}
var RootPageNavbar = []NavbarItem{loginItem, signupItem}

var GoogleItem = OnboardItem{
	Provider: "google",
	Icon:     "google",
}

var GithubItem = OnboardItem{
	Provider: "github",
	Icon:     "github",
}

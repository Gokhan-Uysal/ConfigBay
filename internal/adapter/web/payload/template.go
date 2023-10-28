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

type HomePage struct {
	NavbarItems  []NavbarItem
	ProjectItems []ProjectItem
}

type RootPage struct {
	NavbarItems []NavbarItem
}

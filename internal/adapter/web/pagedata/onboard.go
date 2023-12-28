package pagedata

import "github.com/Gokhan-Uysal/ConfigBay.git/internal/config"

type Access string

const (
	Login  Access = "login"
	Signup        = "signup"
)

type Onboard struct {
	Config *config.OnboardPage
	Access Access
}

package valueobject

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"regexp"
)

type (
	Email interface {
		String() string
	}

	email struct {
		data string
	}
)

const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func MustNewEmail(data string) (Email, error) {
	re, err := regexp.Compile(emailPattern)
	if err != nil {
		return nil, err
	}
	if !re.MatchString(data) {
		return nil, domain.ValidationErr{Info: "email"}
	}

	return &email{data: data}, nil
}

func NewEmail(data string) Email {
	return &email{data: data}
}

func (e *email) String() string {
	return e.data
}

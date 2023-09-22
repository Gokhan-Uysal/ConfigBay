package domain

import (
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

func NewEmail(data string) (Email, error) {
	re, err := regexp.Compile(emailPattern)
	if err != nil {
		return nil, err
	}
	if !re.MatchString(data) {
		return nil, ValidationErr{Item: "email"}
	}

	return &email{data: data}, nil
}

func (e *email) String() string {
	return e.data
}

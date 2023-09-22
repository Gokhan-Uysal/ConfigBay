package domain

import "fmt"

type ValidationErr struct {
	Item string
}

func (e ValidationErr) Error() string {
	return fmt.Sprintf("invalid validation: %s", e.Item)
}

package domain

import "fmt"

type ValidationErr struct {
	Info string
}

type ItemNotFoundErr struct {
	Item string
}

func (e ValidationErr) Error() string {
	return fmt.Sprintf("invalid validation: %s", e.Info)
}

func (e ItemNotFoundErr) Error() string {
	return fmt.Sprintf("item not found in db: %s", e.Item)
}

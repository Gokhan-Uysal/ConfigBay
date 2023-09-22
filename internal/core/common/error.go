package common

import "fmt"

type NilPointerErr struct {
	Item string
}

func (e NilPointerErr) Error() string {
	return fmt.Sprintf("nil pointer at: %s", e.Item)
}

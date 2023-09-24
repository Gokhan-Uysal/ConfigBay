package error

import "fmt"

type (
	UserNotFoundErr struct {
		Field string
	}
)

func (e UserNotFoundErr) Error() string {
	return fmt.Sprintf("user not found with %s", e.Field)
}

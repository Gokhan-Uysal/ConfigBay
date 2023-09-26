package errorx

import "fmt"

type (
	UserNotFoundErr struct {
		Field string
	}

	PermissionErr struct {
		Role string
	}
)

func (e UserNotFoundErr) Error() string {
	return fmt.Sprintf("user not found with %s", e.Field)
}

func (e PermissionErr) Error() string {
	return fmt.Sprintf("permission needed: %s", e.Role)
}

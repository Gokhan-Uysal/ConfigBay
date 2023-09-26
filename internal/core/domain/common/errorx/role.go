package errorx

import "fmt"

type (
	RoleParsingErr struct {
		Name string
	}
)

func (e RoleParsingErr) Error() string {
	return fmt.Sprintf("failed to parse role %s", e.Name)
}

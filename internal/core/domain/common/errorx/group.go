package errorx

import "fmt"

type (
	GroupCreationErr struct {
		Title string
	}
)

func (e GroupCreationErr) Error() string {
	return fmt.Sprintf("group creation failed: %s", e.Title)
}

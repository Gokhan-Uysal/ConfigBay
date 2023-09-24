package error

import "fmt"

type (
	ProjectNotFoundErr struct {
		Id string
	}

	ProjectCreationErr struct {
		Title string
	}
)

func (e ProjectNotFoundErr) Error() string {
	return fmt.Sprintf("project not found with %s", e.Id)
}

func (e ProjectCreationErr) Error() string {
	return fmt.Sprintf("project creation failed: %s", e.Title)
}

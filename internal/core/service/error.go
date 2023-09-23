package service

import "fmt"

type (
	UserNotFoundErr struct {
		Field string
	}

	ProjectCreationErr struct {
		Title string
	}
)

func (e UserNotFoundErr) Error() string {
	return fmt.Sprintf("user not found with %s", e.Field)
}

func (e ProjectCreationErr) Error() string {
	return fmt.Sprintf("project creation failed: %s", e.Title)
}

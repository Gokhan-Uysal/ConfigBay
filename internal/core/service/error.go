package service

import "fmt"

type (
	ProjectNotFoundErr struct {
		Id string
	}

	UserNotFoundErr struct {
		Field string
	}

	ProjectCreationErr struct {
		Title string
	}

	GroupCreationErr struct {
		Title string
	}
)

func (e ProjectNotFoundErr) Error() string {
	return fmt.Sprintf("project not found with %s", e.Id)
}

func (e UserNotFoundErr) Error() string {
	return fmt.Sprintf("user not found with %s", e.Field)
}

func (e ProjectCreationErr) Error() string {
	return fmt.Sprintf("project creation failed: %s", e.Title)
}

func (e GroupCreationErr) Error() string {
	return fmt.Sprintf("group creation failed: %s", e.Title)
}

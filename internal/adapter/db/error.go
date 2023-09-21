package db

import "fmt"

type ItemNotFoundErr struct {
	Item string
}

type UniqueKeyValidationErr struct {
	Key string
}

type QueryError struct {
	Reason string
}

func (e ItemNotFoundErr) Error() string {
	return fmt.Sprintf("item not found: %s", e.Item)
}

func (e UniqueKeyValidationErr) Error() string {
	return fmt.Sprintf("unique key validation: %s", e.Key)
}

func (e QueryError) Error() string {
	return fmt.Sprintf("query failed: %s", e.Reason)
}

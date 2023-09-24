package db

import "fmt"

type ItemNotFoundErr struct {
	Item string
}

type UniqueKeyValidationErr struct {
	Key string
}

type MappingErr struct {
	Item string
}

func (e ItemNotFoundErr) Error() string {
	return fmt.Sprintf("item not found: %s", e.Item)
}

func (e UniqueKeyValidationErr) Error() string {
	return fmt.Sprintf("unique key validation: %s", e.Key)
}

func (e MappingErr) Error() string {
	return fmt.Sprintf("mapping failed: %s", e.Item)
}

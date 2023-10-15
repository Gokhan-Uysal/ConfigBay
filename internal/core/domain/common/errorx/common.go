package errorx

import "fmt"

type (
	ValidationErr struct {
		Info string
	}

	ItemNotFoundErr struct {
		Item string
	}

	NilPointerErr struct {
		Item string
	}
)

func (e ValidationErr) Error() string {
	return fmt.Sprintf("invalid validation: %s", e.Info)
}

func (e ItemNotFoundErr) Error() string {
	return fmt.Sprintf("item not found in db: %s", e.Item)
}

func (e NilPointerErr) Error() string {
	return fmt.Sprintf("nil pointer at: %s", e.Item)
}

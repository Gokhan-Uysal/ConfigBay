package valueobject

import "github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/model"

type ProjectID interface {
	model.ID
}

type GroupID interface {
	model.ID
}

type UserID interface {
	model.ID
}

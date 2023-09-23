package port

import "github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"

type (
	GroupRepo interface {
		Save(group aggregate.Group) error
	}
)

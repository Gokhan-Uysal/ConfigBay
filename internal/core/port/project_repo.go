package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
)

type (
	ProjectRepo interface {
		Save(project aggregate.Project) error
	}
)

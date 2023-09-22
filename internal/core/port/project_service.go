package port

import "github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"

type (
	ProjectService interface {
		Init(
			userId domain.ID,
			projectTitle string,
			groupTitle string,
		) (domain.Project,
			error)
	}
)

package port

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/adapter/web/payload"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
)

type (
	ProjectService interface {
		Create(
			valueobject.UserID,
			string,
			string,
		) (aggregate.Project,
			error)
		Find(
			valueobject.ProjectID,
			valueobject.UserID,
		) (aggregate.Project, error)
	}

	GroupService interface {
		CreateGroup(
			groupTitle string,
			projectId valueobject.ProjectID,
			role valueobject.Role,
			userIds ...valueobject.UserID,
		) (aggregate.Group, error)
	}

	UserService interface {
		Find(valueobject.UserID) (aggregate.User, error)
	}

	GoogleAuthService interface {
		BuildSSO(provider string) payload.SSO
		FetchToken(code string) (*payload.GoogleToken, error)
		RefreshToken(refreshToken string) (*payload.GoogleToken, error)
		RevokeToken(token string) error
	}
)

package pagedata

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/config"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/aggregate"
	"time"
)

type HomePage struct {
	Config       *config.HomePage
	ProjectItems []ProjectItem
}

type ProjectItem struct {
	Icon        []byte
	Title       string
	Description string
	GroupCount  int
	UserCount   int
	CreatedAt   time.Time
}

func ToProjectItem(project aggregate.Project) ProjectItem {
	return ProjectItem{
		Icon:        nil,
		Title:       project.Title(),
		Description: "",
		GroupCount:  len(project.Groups()),
		UserCount:   0,
		CreatedAt:   project.CreatedAt(),
	}
}

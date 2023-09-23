package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"time"
)

type (
	ProjectBuilder interface {
		Secrets(...entity.Secret) ProjectBuilder
		Groups(...valueobject.ID) ProjectBuilder
		CreatedAt(time.Time) ProjectBuilder
		UpdatedAt(time.Time) ProjectBuilder
		domain.Builder[Project]
	}

	Project interface {
		BaseAggregate
		Title() string
		Secrets() []entity.Secret
		Groups() []valueobject.ID
		AddGroup(group valueobject.ID)
		AddSecret(secret entity.Secret)
	}

	projectBuilder struct {
		project
	}

	project struct {
		*baseAggregate
		title   string
		secrets []entity.Secret
		groups  []valueobject.ID
	}
)

func NewProjectBuilder(id valueobject.ID, title string) ProjectBuilder {
	base := newBaseAggregate(id)
	return &projectBuilder{project: project{baseAggregate: base, title: title}}
}

func (pb *projectBuilder) Secrets(secrets ...entity.Secret) ProjectBuilder {
	pb.secrets = secrets
	return pb
}

func (pb *projectBuilder) Groups(groups ...valueobject.ID) ProjectBuilder {
	pb.groups = groups
	return pb
}

func (pb *projectBuilder) CreatedAt(time time.Time) ProjectBuilder {
	pb.createdAt = time
	return pb
}

func (pb *projectBuilder) UpdatedAt(time time.Time) ProjectBuilder {
	pb.updatedAt = time
	return pb
}

func (pb *projectBuilder) Build() Project {
	return &project{
		baseAggregate: pb.baseAggregate,
		title:         pb.title,
		groups:        pb.groups,
		secrets:       pb.secrets,
	}
}

func (p *project) Title() string {
	return p.title
}

func (p *project) Groups() []valueobject.ID {
	return p.groups
}

func (p *project) Secrets() []entity.Secret {
	return p.secrets
}

func (p *project) AddGroup(group valueobject.ID) {
	p.groups = append(p.groups, group)
}

func (p *project) AddSecret(secret entity.Secret) {
	p.secrets = append(p.secrets, secret)
}

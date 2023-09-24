package aggregate

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/model"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/entity"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	valueobject2 "github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"time"
)

type (
	ProjectBuilder interface {
		Secrets(...entity.Secret) ProjectBuilder
		Groups(...valueobject.GroupID) ProjectBuilder
		CreatedAt(time.Time) ProjectBuilder
		UpdatedAt(time.Time) ProjectBuilder
		model.Builder[Project]
	}

	Project interface {
		model.BaseAggregate
		Title() string
		Secrets() []entity.Secret
		Groups() []valueobject.GroupID
		AddGroup(group valueobject.GroupID)
		AddGroups(groups ...valueobject.GroupID)
		AddSecret(secret entity.Secret)
		AddSecrets(secrets ...entity.Secret)
	}

	projectBuilder struct {
		project
	}

	project struct {
		*baseAggregate
		title   string
		secrets []entity.Secret
		groups  []valueobject.GroupID
	}
)

func NewProjectBuilder(id valueobject2.ProjectID, title string) ProjectBuilder {
	base := newBaseAggregate(id)
	return &projectBuilder{project: project{baseAggregate: base, title: title}}
}

func (pb *projectBuilder) Secrets(secrets ...entity.Secret) ProjectBuilder {
	pb.secrets = secrets
	return pb
}

func (pb *projectBuilder) Groups(groups ...valueobject.GroupID) ProjectBuilder {
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

func (p *project) Groups() []valueobject.GroupID {
	return p.groups
}

func (p *project) Secrets() []entity.Secret {
	return p.secrets
}

func (p *project) AddGroup(group valueobject.GroupID) {
	p.groups = append(p.groups, group)
}

func (p *project) AddGroups(groups ...valueobject.GroupID) {
	p.groups = append(p.groups, groups...)
}

func (p *project) AddSecret(secret entity.Secret) {
	p.secrets = append(p.secrets, secret)
}

func (p *project) AddSecrets(secrets ...entity.Secret) {
	p.secrets = append(p.secrets, secrets...)
}

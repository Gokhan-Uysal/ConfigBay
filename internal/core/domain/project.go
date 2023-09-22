package domain

import "time"

type (
	ProjectBuilder interface {
		Id(int) ProjectBuilder
		Secrets(...Secret) ProjectBuilder
		Groups(...Group) ProjectBuilder
		CreatedAt(time.Time) ProjectBuilder
		UpdatedAt(time.Time) ProjectBuilder
		Builder[Project]
	}

	Project interface {
		Id() int
		Title() string
		Secrets() []Secret
		Groups() []Group
		AddGroup(group Group)
		AddSecret(secret Secret)
		Timestamp
	}

	projectBuilder struct {
		project
	}

	project struct {
		id        int
		title     string
		secrets   []Secret
		groups    []Group
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewProjectBuilder(title string) ProjectBuilder {
	return &projectBuilder{project{id: -1, title: title}}
}

func (pb *projectBuilder) Id(id int) ProjectBuilder {
	pb.id = id
	return pb
}

func (pb *projectBuilder) Secrets(secrets ...Secret) ProjectBuilder {
	pb.secrets = secrets
	return pb
}

func (pb *projectBuilder) Groups(groups ...Group) ProjectBuilder {
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
		id:        pb.id,
		title:     pb.title,
		groups:    pb.groups,
		secrets:   pb.secrets,
		createdAt: pb.createdAt,
		updatedAt: pb.updatedAt,
	}
}

func (p *project) Id() int {
	return p.id
}

func (p *project) Title() string {
	return p.title
}

func (p *project) Groups() []Group {
	return p.groups
}

func (p *project) Secrets() []Secret {
	return p.secrets
}

func (p *project) CreatedAt() time.Time {
	return p.createdAt
}

func (p *project) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *project) AddGroup(group Group) {
	p.groups = append(p.groups, group)
}

func (p *project) AddSecret(secret Secret) {
	p.secrets = append(p.secrets, secret)
}

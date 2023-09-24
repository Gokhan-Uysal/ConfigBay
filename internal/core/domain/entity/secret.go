package entity

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/valueobject"
	"time"
)

type (
	SecretBuilder interface {
		Key(string) SecretBuilder
		Value(string) SecretBuilder
		Version(int) SecretBuilder
		CreatedAt(time.Time) SecretBuilder
		UpdatedAt(time.Time) SecretBuilder
		domain.Builder[Secret]
	}

	Secret interface {
		Id() valueobject.ID
		Key() string
		Value() string
		CreatedAt() time.Time
		UpdatedAt() time.Time
	}

	secretBuilder struct {
		secret
	}

	secret struct {
		id        valueobject.ID
		key       string
		value     string
		version   int
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewSecretBuilder(id valueobject.ID) SecretBuilder {
	return &secretBuilder{secret{id: id, version: 1}}
}

func (sb *secretBuilder) Key(k string) SecretBuilder {
	sb.key = k
	return sb
}

func (sb *secretBuilder) Value(v string) SecretBuilder {
	sb.value = v
	return sb
}

func (sb *secretBuilder) Version(v int) SecretBuilder {
	sb.version = v
	return sb
}

func (sb *secretBuilder) CreatedAt(t time.Time) SecretBuilder {
	sb.createdAt = t
	return sb
}

func (sb *secretBuilder) UpdatedAt(t time.Time) SecretBuilder {
	sb.updatedAt = t
	return sb
}

func (sb *secretBuilder) Build() Secret {
	return &secret{
		id:        sb.id,
		key:       sb.key,
		value:     sb.value,
		version:   sb.version,
		createdAt: sb.createdAt,
		updatedAt: sb.updatedAt,
	}
}

func (s *secret) Id() valueobject.ID {
	return s.id
}

func (s *secret) Key() string {
	return s.key
}

func (s *secret) Value() string {
	return s.value
}

func (s *secret) CreatedAt() time.Time {
	return s.createdAt
}

func (s *secret) UpdatedAt() time.Time {
	return s.updatedAt
}

package entity

import (
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain"
	"time"
)

type (
	SecretBuilder interface {
		Version(int) SecretBuilder
		CreatedAt(time.Time) SecretBuilder
		UpdatedAt(time.Time) SecretBuilder
		domain.Builder[Secret]
	}

	Secret interface {
		Key() string
		Value() string
		CreatedAt() time.Time
		UpdatedAt() time.Time
	}

	secretBuilder struct {
		secret
	}

	secret struct {
		key       string
		value     string
		version   int
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewSecretBuilder(key, value string) SecretBuilder {
	return &secretBuilder{secret{key: key, value: value, version: 1}}
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
		key:       sb.key,
		value:     sb.value,
		version:   sb.version,
		createdAt: sb.createdAt,
		updatedAt: sb.updatedAt,
	}
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

package domain

import "time"

type (
	SecretBuilder interface {
		Key(string) SecretBuilder
		Value(string) SecretBuilder
		CreatedAt(time.Time) SecretBuilder
		UpdatedAt(time.Time) SecretBuilder
		Builder[Secret]
	}

	Secret interface {
		Id() ID
		Key() string
		Value() string
		Timestamp
	}

	secretBuilder struct {
		secret
	}

	secret struct {
		id        ID
		key       string
		value     string
		createdAt time.Time
		updatedAt time.Time
	}
)

func NewSecretBuilder(id ID, key string, value string) SecretBuilder {
	return &secretBuilder{secret{id: id, key: key, value: value}}
}

func (sb *secretBuilder) Key(key string) SecretBuilder {
	sb.key = key
	return sb
}

func (sb *secretBuilder) Value(value string) SecretBuilder {
	sb.value = value
	return sb
}

func (sb *secretBuilder) CreatedAt(time time.Time) SecretBuilder {
	sb.createdAt = time
	return sb
}

func (sb *secretBuilder) UpdatedAt(time time.Time) SecretBuilder {
	sb.updatedAt = time
	return sb
}

func (sb *secretBuilder) Build() Secret {
	return &secret{
		id:        sb.id,
		key:       sb.key,
		value:     sb.value,
		createdAt: sb.createdAt,
		updatedAt: sb.updatedAt,
	}
}

func (s *secret) Id() ID {
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
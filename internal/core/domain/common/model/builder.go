package model

type Builder[Entity interface{}] interface {
	Build() Entity
}

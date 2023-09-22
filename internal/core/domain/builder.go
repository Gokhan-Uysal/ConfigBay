package domain

type Builder[Entity interface{}] interface {
	Build() Entity
}

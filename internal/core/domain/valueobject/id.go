package valueobject

type (
	ID interface {
		String() string
	}

	ProjectID interface {
		ID
	}

	GroupID interface {
		ID
	}

	UserID interface {
		ID
	}
)

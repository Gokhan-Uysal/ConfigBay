package domain

import "time"

type Timestamp interface {
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

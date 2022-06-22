package product

import "time"

type Product struct {
	ID          int64
	ProductName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Deleted     bool
}

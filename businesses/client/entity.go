package client

import "time"

type Client struct {
	ID          int64
	ClientID    string
	ClientAlias string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Deleted     bool
}

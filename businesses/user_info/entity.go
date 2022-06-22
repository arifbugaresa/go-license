package user_info

import "time"

type UserInfo struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
	Deleted   time.Time
}

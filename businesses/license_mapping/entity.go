package license_mapping

import "time"

type LicenseMapping struct {
	ID        int64
	ClientID  int64
	ProductID int64
	LicenseID int64
	Deleted   bool
	CreatedAt time.Time
	DeletedAt time.Time
}

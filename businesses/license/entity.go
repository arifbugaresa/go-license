package license

import "time"

type License struct {
	ID               int64
	LicenseValidFrom time.Time
	LicenseValidThru time.Time
	CreatedAt        time.Time
	DeletedAt        time.Time
	Deleted          bool
}

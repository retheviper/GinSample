package entity

import "time"

type Audit struct {
	CreatedBy        string
	CreatedDate      time.Time
	LastModifiedBy   string
	LastModifiedDate time.Time
}

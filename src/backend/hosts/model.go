package hosts

import (
	"time"
)

type Host struct {
	ID 			string		`json:"id" db:"id"`
	Url			string		`json:"url" db:"url"`
	CreatedAt	time.Time	`json:"createdAt" db:"created_at"`
}

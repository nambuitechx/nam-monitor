package healthlogs

import (
	"time"
)

type HealthLog struct {
	ID 			string		`json:"id" db:"id"`
	HostID		string		`json:"hostID" db:"host_id"`
	Status		string		`json:"status" db:"status"`
	CreatedAt	time.Time	`json:"createdAt" db:"created_at"`
}

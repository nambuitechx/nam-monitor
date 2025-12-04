package healthlogs

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3/table"
)

type HealthLog struct {
	ID 			gocql.UUID	`json:"id" db:"id"`
	Status		string		`json:"status" db:"status"`
	CreatedAt	time.Time	`json:"createdAt" db:"created_at"`
}

var HealthLogTable = table.New(table.Metadata{
	Name: "health_logs",
	Columns: []string{"id", "status", "created_at"},
	PartKey: []string{"id"},
	SortKey: []string{"created_at"},
})

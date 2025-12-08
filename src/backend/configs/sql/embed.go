package sql

import "embed"

//go:embed *.sql
var MigrationFiles embed.FS

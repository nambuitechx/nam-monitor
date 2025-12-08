package cql

import "embed"

//go:embed *.cql
var MigrationFiles embed.FS

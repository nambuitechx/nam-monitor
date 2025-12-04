module github.com/nambuitechx/go-monitor/backend

go 1.25.0

require (
	github.com/go-chi/chi/v5 v5.2.3
	github.com/go-chi/cors v1.2.2
	github.com/gocql/gocql v1.7.0
	github.com/scylladb/gocqlx/v3 v3.0.4
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.18.1 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	golang.org/x/sync v0.17.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

// Use the latest version of scylladb/gocql; check for updates at https://github.com/scylladb/gocql/releases
replace github.com/gocql/gocql => github.com/scylladb/gocql v1.16.0

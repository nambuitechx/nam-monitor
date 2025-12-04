package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/migrate"
	mycql "github.com/nambuitechx/go-monitor/backend/configs/cql"
)

type DBConnection struct {
	Session *gocqlx.Session
}

var Session *gocqlx.Session

func NewDBConnection(envConfig *EnvConfig) *DBConnection {
	// Create gocql cluster.
	hosts := []string {envConfig.DBHost}
	cluster := gocql.NewCluster(hosts...)
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 5 * time.Second

	// Init session
	log.Printf("Setting up database connection to %s...\n", envConfig.DBHost)
	var sysSession *gocql.Session
	var err error
	retries := 1

	for {
		if retries > 10 {
			log.Fatalf("Connect to database failed: %s", err)
		}

		log.Printf("Trying to connect to %s %d times\n", envConfig.DBHost, retries)
		sysSession, err = cluster.CreateSession()

		if err != nil {
			time.Sleep(5 * time.Second)
			retries += 1
		} else {
			break
		}
	}

	log.Println("Database connected!")

	// Create keyspace if not exists
    cql := fmt.Sprintf(`
    CREATE KEYSPACE IF NOT EXISTS %s
    WITH replication = {
        'class': 'NetworkTopologyStrategy',
        'datacenter1': 1
    }
	AND tablets = {'enabled': false};`, envConfig.DBName)

	if err := sysSession.Query(cql).Exec(); err != nil {
		sysSession.Close()
		log.Fatal(err)
	}

	sysSession.Close()

	cluster.Keyspace = envConfig.DBName
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}

	err = migrate.FromFS(context.Background(), session, mycql.Files)
	if err != nil {
		log.Fatalf("Run database migration failed: %s", err)
	}

	log.Println("Run database migration successfully!")
	Session = &session

	return &DBConnection{
		Session: &session,
	}
}

func CloseDBConnection() {
	log.Println("Closing DB connection...")

	if !Session.Closed() {
		Session.Close()
	}
	
	log.Println("DB connection closed!")
}

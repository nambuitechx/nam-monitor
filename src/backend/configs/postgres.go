package configs

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	mysql "github.com/nambuitechx/go-monitor/backend/configs/sql"
)

type PostgresConnection struct {
	DB *sqlx.DB
}

var DB *sqlx.DB

func NewPostgresConnection(envConfig *EnvConfig) *PostgresConnection {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		envConfig.DBUser,
		envConfig.DBPassword,
		envConfig.DBHost,
		envConfig.DBPort,
		envConfig.DBName,
	)

	log.Printf("Setting up Postgres database connection to %s:%s...\n", envConfig.DBHost, envConfig.DBPort)
	var db *sql.DB
	var err error
	retries := 1

	for {
		if retries > 5 {
			log.Fatalf("Connect to Postgres database failed: %s", err)
		}

		log.Printf("Trying to connect to Postgres: %s:%s %d times\n", envConfig.DBHost, envConfig.DBPort, retries)
		db, err = sql.Open("postgres", dsn)

		if err != nil {
			time.Sleep(5 * time.Second)
			retries += 1
		} else {
			err = db.Ping()

			if err != nil {
				time.Sleep(5 * time.Second)
				retries += 1
			} else {
				break
			}
		}
	}

	source, _ := iofs.New(mysql.MigrationFiles, ".")
	m, err := migrate.NewWithSourceInstance("iofs", source, dsn)

    if err != nil {
		log.Fatalf("Failed to migrate instance: %s", err)
    }

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migration: %v", err)
	}

	log.Println("Database migrations run successfully")

	sqlxDB := sqlx.NewDb(db, "postgres")
	DB = sqlxDB

	return &PostgresConnection{DB: sqlxDB}
}

func ClosePostgresConnection() {
	log.Println("Closing Postgres connection...")

	if err := DB.Close(); err != nil {
		log.Fatalln("Close Postgres connection failed")
	}
	
	log.Println("Postgres connection closed!")
}

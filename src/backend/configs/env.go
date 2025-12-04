package configs

import "os"

type EnvConfig struct {
	Host		string
	Port		string

	DBHost 		string
	DBPort 		string
	DBName 		string
	DBUser 		string
	DBPassword	string
}

func NewEnvConfig() *EnvConfig {
	host := getEnv("HOST", "localhost")
	port := getEnv("PORT", "8000")

	dbHost := getEnv("DB_HOST", "scylla")
	dbPort := getEnv("DB_PORT", "9042")
	dbName := getEnv("DB_NAME", "demo")
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "admin")

	return &EnvConfig{
		Host: host,
		Port: port,
		DBHost: dbHost,
		DBPort: dbPort,
		DBName: dbName,
		DBUser: dbUser,
		DBPassword: dbPassword,
	}
}

func getEnv(k string, d string) string {
	v := os.Getenv(k)

	if v != "" {
		return v
	}

	return d
}

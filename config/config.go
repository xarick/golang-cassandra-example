package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Application struct {
	GinMode string
	RunPort string

	CassandraHost string
	CassandraKey  string
}

func LoadConfig() Application {
	cfg := Application{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	cfg.GinMode = os.Getenv("GIN_MODE")
	cfg.RunPort = os.Getenv("RUN_PORT")

	cfg.CassandraHost = os.Getenv("CASSANDRA_HOST")
	cfg.CassandraKey = os.Getenv("CASSANDRA_KEY")

	return cfg
}

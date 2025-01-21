package main

import (
	"log"

	"github.com/xarick/golang-cassandra-example/config"
	"github.com/xarick/golang-cassandra-example/db"
	"github.com/xarick/golang-cassandra-example/routes"
)

func main() {
	cfg := config.LoadConfig()

	db.InitCassandra(cfg)
	defer db.CloseCassandra()

	r := routes.SetupRoutes()

	if err := r.Run(cfg.RunPort); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}

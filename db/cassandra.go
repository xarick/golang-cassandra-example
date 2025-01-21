package db

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/xarick/golang-cassandra-example/config"
)

var Session *gocql.Session

func InitCassandra(cfg config.Application) {
	cluster := gocql.NewCluster(cfg.CassandraHost)
	// cluster.Port = 9042 // Cassandra standart porti
	cluster.Keyspace = cfg.CassandraKey
	cluster.Consistency = gocql.Quorum

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Cassandra bilan bog'lanishda xatolik: %v", err)
	}
	log.Println("Cassandra muvaffaqiyatli ulanadi.")
}

func CloseCassandra() {
	if Session != nil {
		Session.Close()
	}
}

package main

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"timescaledb-benchmark-assignment/infra/db/timescaledb"
)

const (
	migrationsSourceUrl = "file://migrations"
)

func main() {
	log.Info("Starting database migration")

	db, err := timescaledb.GetDb()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	migration := NewDbMigration(db)

	if err = migration.Run(migrationsSourceUrl); err != nil {
		log.Fatalf("Unable to run database migrations: %s", err)
	}

	log.Info("Database migration completed successfully")
}

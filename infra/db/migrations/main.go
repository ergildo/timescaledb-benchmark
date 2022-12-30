package main

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"timescaledb-benchmark-assignment/infra/db/timescaledb"
)

var (
	migrationsSourceUrl = os.Getenv("MIGRATION_SOURCE_URL")
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

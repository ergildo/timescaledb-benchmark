package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

type DbMigration interface {
	Run(sourceUrl string) error
}

type dbMigrationImpl struct {
	db *sql.DB
}

// NewDbMigration create DbMigration
func NewDbMigration(db *sql.DB) DbMigration {
	return dbMigrationImpl{
		db: db,
	}
}

// Run execute database migrations
func (t dbMigrationImpl) Run(sourceUrl string) error {
	log.Info("Running migration scripts")
	driver, err := postgres.WithInstance(t.db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		sourceUrl,
		"postgres", driver)

	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

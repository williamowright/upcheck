package migrations

import (
    _ "github.com/lib/pq"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)


func Migrate(dbConnStr string) {
    m, err := migrate.New(
			"file://migrations",
			dbConnStr)
	if err != nil {
		panic(err)
	}
	migrationStatus := m.Up()
	if migrationStatus != nil && migrationStatus != migrate.ErrNoChange {
		panic(migrationStatus)
	}
}

package sqlserver

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	mssqlserver "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"profile/internal/cfg"
)

func Start(config *cfg.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		config.SqlServerConfig.User, config.SqlServerConfig.Password, config.SqlServerConfig.Host, config.SqlServerConfig.Port, config.SqlServerConfig.Database,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func RunMigrations(config *cfg.Config) error {
	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		config.SqlServerConfig.User, config.SqlServerConfig.Password, config.SqlServerConfig.Host, config.SqlServerConfig.Port, config.SqlServerConfig.Database,
	)

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return err
	}

	driver, err := mssqlserver.WithInstance(db, &mssqlserver.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://platform/migrations",
		"profile", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

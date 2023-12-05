package sqlserver

import (
	"fmt"
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

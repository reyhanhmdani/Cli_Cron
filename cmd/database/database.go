package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigration "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"pr_ramadhan/cmd/models"
	"time"
)

func ConnnectDb(cfg *models.Config) (*gorm.DB, error) {

	fmt.Printf("%+v\n", cfg)
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		panic("Cannot Connect to Database")
	}

	logrus.Info("Success Connect to Database")
	return db, err
}

func Migrate(db *gorm.DB) error {
	logrus.Info("running database migration")

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	driver, err := mysqlMigration.WithInstance(sqlDB, &mysqlMigration.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"mysql", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err == migrate.ErrNoChange {
		logrus.Info("No schema changes to apply")
		return nil
	}

	return err
}

package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	models2 "pr_ramadhan/cmd/models"
	"time"
)

func ConnnectDb(cfg *models2.Config) (*gorm.DB, error) {

	fmt.Printf("%+v\n", cfg)
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
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

	err = db.AutoMigrate(&models2.Wikis{})
	if err != nil {
		return nil, err
	}

	logrus.Info("Success Connect to Database")
	return db, err
}

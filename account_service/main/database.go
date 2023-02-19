package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"store-bpel/account_service/config"
)

func dsn(dbConfig *config.MySQLConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
}

func DbConnect(dbConfig *config.MySQLConfig) (*gorm.DB, error) {
	dbDsn := dsn(dbConfig)
	db, err := gorm.Open(mysql.Open(dbDsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// test db connection
	err = db.Raw("SELECT 1").Error
	if err != nil {
		return nil, err
	}

	return db, nil
}

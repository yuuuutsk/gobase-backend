package database

import (
	"time"

	"github.com/go-sql-driver/mysql"

	"database/sql"
)

func NewDB(config *DBConfig) (*sql.DB, error) {

	net := "tcp"
	if config.DbNet != "" {
		net = config.DbNet
	}

	c := mysql.Config{
		DBName:               config.DbName,
		User:                 config.DbUser,
		Passwd:               config.DbPassword,
		Addr:                 config.DbAddr,
		Net:                  net,
		ParseTime:            true,
		AllowNativePasswords: true,
		Loc:                  time.FixedZone("Asia/Tokyo", 9*60*60),
	}
	dsn := c.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type DBConfig struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbAddr     string
	DbNet      string
}

func NewDBConfig(
	dbName string,
	dbUser string,
	dbPassword string,
	DbAddr string,
	DbNet string,
) *DBConfig {
	return &DBConfig{
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbAddr:     DbAddr,
		DbNet:      DbNet,
	}
}

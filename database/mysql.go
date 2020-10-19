package database

import (
	"fmt"
	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	cfg, err := ini.Load("./config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
	}
	host := cfg.Section("mysql").Key("DB_HOST").String()
	port := cfg.Section("mysql").Key("DB_PORT").String()
	database := cfg.Section("mysql").Key("DB_DATABASE").String()
	user := cfg.Section("mysql").Key("DB_USER").String()
	password := cfg.Section("mysql").Key("DB_PASSWORD").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Printf("DB:%s", err)
	}
}

func CloseDB() {
	defer db.Close()
}

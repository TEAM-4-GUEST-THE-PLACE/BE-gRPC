package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )
  
  var DB *gorm.DB
  
  func init() {
	DatabaseConnection()
  }
  
  func DatabaseConnection() {
	dsn := "host=roundhouse.proxy.rlwy.net port=32190 user=postgres dbname=railway password=kqaUZaPvstRpAiHSUDeBbVxVPtkgpEhv sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	DB = db
  }
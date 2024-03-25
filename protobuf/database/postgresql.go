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
	dsn := "host=postgresql-167179-0.cloudclusters.net port=19727 user=admin dbname=be-GTP password=admin123 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	DB = db
  }
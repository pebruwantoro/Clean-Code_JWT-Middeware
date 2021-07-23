package config

import (
	"fmt"
	"pebruwantoro/middleware/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		//you can user your own account SQL
		"DB_Username": "",
		"DB_Password": "",
		"DB_Host":     "",
		"DB_Port":     "",
		"DB_Name":     "",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

//Automigration to database SQL
func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}

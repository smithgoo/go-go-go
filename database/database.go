package database

import "gorm.io/driver/mysql"
import "gorm.io/gorm"
import "log"

var DB *gorm.DB

func InitDB(){
	dsn := "root:Jbnb123456@tcp(127.0.0.1:3306)/GODB?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err !=nil {
		log.Fatal("Failed to connect to database:",err)
	}
	log.Println("Database connected successfully!")
}
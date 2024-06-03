package database

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"strconv"
)
import "gorm.io/gorm"
import "log"

var DB *gorm.DB

func InitDB(user,password,host,dbName string){
	//dsn := "root:Jbnb123456@tcp(127.0.0.1:3306)/GODB?charset=utf8&parseTime=True&loc=Local"

	dsn := user + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	println("DSN: ", dsn)  // For debugging purposes
	var err error
	DB,err = gorm.Open(mysql.Open(dsn),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err !=nil {
		log.Fatal("Failed to connect to database:",err)
	}
	log.Println("Database connected successfully!")


	//var err error
	//DB,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	//if err !=nil {
	//	log.Fatal("Failed to connect to database:",err)
	//}
	//log.Println("Database connected successfully!")
}

type PaginatedResult struct {
	Page      int         `json:"page"`
	Size  int         `json:"size"`
	Total     int64       `json:"total"`
	Data      interface{} `json:"data"`
}

func Paginate(c *gin.Context, db *gorm.DB, model interface{}, out interface{}) (PaginatedResult, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		size = 10
	}

	offset := (page - 1) * size

	var total int64
	result := db.Model(model).Count(&total)
	if result.Error != nil {
		return PaginatedResult{}, result.Error
	}

	result = db.Limit(size).Offset(offset).Find(out)
	if result.Error != nil {
		return PaginatedResult{}, result.Error
	}

	return PaginatedResult{
		Page:     page,
		Size: size,
		Total:    total,
		Data:     out,
	}, nil
}


func PaginateSearch(c *gin.Context, db *gorm.DB, model interface{}, out interface{}, query *gorm.DB) (PaginatedResult, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	size, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || size < 1 {
		size = 10
	}

	offset := (page - 1) * size

	var total int64
	result := query.Model(model).Count(&total)
	if result.Error != nil {
		return PaginatedResult{}, result.Error
	}

	result = query.Limit(size).Offset(offset).Find(out)
	if result.Error != nil {
		return PaginatedResult{}, result.Error
	}

	return PaginatedResult{
		Page:     page,
		Size: size,
		Total:    total,
		Data:     out,
	}, nil
}
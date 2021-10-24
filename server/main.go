package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/yoshikawa/agetaro/server/database"
	"github.com/yoshikawa/agetaro/server/models"
)

func isCloudRun() bool {
	return os.Getenv("PROJECT_ID") != ""
}

func main() {
	var db *gorm.DB
	var err error
	if isCloudRun() {
		db, err = database.GetCloudSQLConn()
	} else {
		db, err = database.GetGormConn()
	}

	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		user := models.User{
			ID:       uuid.New(),
			Name:     "hoge",
			Email:    "hoge@example.com",
			Password: "hogehoge",
		}
		db.Create(&user)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

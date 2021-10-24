package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			"https://agetaro-2517d.web.app/",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

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

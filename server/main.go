package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/yoshikawa/agetaro/server/database"
	"github.com/yoshikawa/agetaro/server/models"
	"golang.org/x/crypto/bcrypt"
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

	// session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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
			"https://agetaro-2517d.web.app",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.POST("/signup", func(c *gin.Context) {
		user := models.User{}
		err := c.Bind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request")
			return
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request")
			return
		}
		user.Password = string(hashed)
		user.ID = uuid.New()
		db.NewRecord(user)
		db.Create(&user)
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		user := models.User{}
		du := models.User{}
		err := c.Bind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request")
			return
		}
		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if result := db.Where("email = ?", user.Email).First(&du); result.Error != nil {
			c.JSON(http.StatusBadRequest, "Failed to get data")
			return
		}

		err = bcrypt.CompareHashAndPassword(hashed, []byte(du.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request")
			return
		}
		user.ID = uuid.New()
		session := sessions.Default(c)
		session.Set("loginUser", du.ID)
		session.Save()
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

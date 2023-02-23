package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ydammatsu/toybox/database"
	"github.com/ydammatsu/toybox/model"
	"gorm.io/gorm"
)

func dbInit() *gorm.DB {
	db := database.Connect()
	db.AutoMigrate(&model.User{})
	return db
}

func main() {
	db := dbInit()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		u := &model.User{}
		db.First(u)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!" + u.Name,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

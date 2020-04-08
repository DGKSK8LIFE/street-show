package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dilacts/mysql"
)

var db 

func init() {
	db, err := gorm.Open("mysql", "[dbinfo]")
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Content) {
	
	})
	r.POST("/buskerlist", func(c *gin.Content) {

	})

}

package main

import (
	"street/app/backend/dbcrud"
	"street/app/backend/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	dbcrud.Open("db_info.yaml")
}

func main() {
	r := gin.Default()
	r.GET("/", handler.ServeHome)
	r.Run()
}

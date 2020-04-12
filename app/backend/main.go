package main

import (
	"street-show/app/backend/dbcrud"
	"street-show/app/backend/handler"

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

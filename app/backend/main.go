package main

import (
	"street-show/app/backend/dbcrud"
	"street-show/app/backend/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// init precedingly opens database
func init() {
	dbcrud.Open("db_info.yaml")
}

func main() {
	r := gin.Default()
	r.GET("/api/busker", handler.BuskerApi)
	r.GET("/api/user", handler.UserApi)
	r.Run()
}

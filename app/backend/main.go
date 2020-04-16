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

// main just declares routes and imports the handler module (mine) to serve accordingly
func main() {
	r := gin.Default()
	r.GET("/api/search/busker", handler.BuskerSearchApi)
	r.GET("/api/search/user", handler.UserSearchApi)
	r.Run()
}

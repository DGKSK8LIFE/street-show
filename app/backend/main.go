package main

import (
	"street-show/app/backend/dbcrud"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	dbcrud.Open("db_info.yaml")
}

func main() {
	r := gin.Default()
	r.Run()
}

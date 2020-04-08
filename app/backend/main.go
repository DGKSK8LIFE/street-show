package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dilacts/mysql"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type DB_info struct {
	User string  `yaml:"user"`
	Password string `yaml:"password"`
	Port uint32 	`yaml:"port"`
	Db_name string 	`yaml:"db_name"`
}

var db 

func init() {
	infoStruct := DB_info{}
	file, err := ioutil.ReadFile("db_info.yaml")
	if err != nil {
		log.Fatalf("database info file error: %s\n", err)	
	}
	err = yaml.Unmarshal(file, infoStruct)
	db, err := gorm.Open("mysql", "[dbinfo]")
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Content) {
	
	})
	r.POST("/buskerlist", func(c *gin.Content) {

	})

}

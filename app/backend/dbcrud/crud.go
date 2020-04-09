package dbcrud

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
)

type DB_info struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     uint32 `yaml:"port"`
	Db_name  string `yaml:"db_name"`
}

var db *gorm.DB

func Open(filename string) {
	infoStruct := &DB_info{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("database info file error: %s\n", err)
	}
	err = yaml.Unmarshal(file, infoStruct)
	if err != nil {
		log.Fatalf("unmarshalling problem: %s\n", err)
	}
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", infoStruct.User, infoStruct.Password, infoStruct.Port, infoStruct.Db_name))
	if err != nil {
		log.Fatalf("database opening error: %s\n", err)
	}
}

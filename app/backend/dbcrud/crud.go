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

type UserAndBusker struct {
	Username string
	Name     string
	Email    string
	Id       uint64
}

var DB *gorm.DB

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
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", infoStruct.User, infoStruct.Password, infoStruct.Port, infoStruct.Db_name))
	if err != nil {
		log.Fatalf("database opening error: %s\n", err)
	}
}

func SelectAll(u *UserAndBusker) {
	DB.Find(&u)
}

func SelectLike(u *UserAndBusker, likeString string) {
	DB.Where("username LIKE ? OR name LIKE ?", likeString, likeString).Find(&u)
}

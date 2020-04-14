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
}

type Busker struct {
	Username string `gorm"column:username" json:"username"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Id       uint64 `gorm:"column:id" json:"id"`
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
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/Street-Show", infoStruct.User, infoStruct.Password, infoStruct.Port))
	if err != nil {
		log.Fatalf("database opening error: %s\n", err)
	}
}

func (u *Busker) SelectAll() {
	DB.Find(&u)
}

func (u *Busker) SelectLike(likeString string) {
	DB.Where("username LIKE ? OR name LIKE ?", likeString, likeString).Find(&u)
}

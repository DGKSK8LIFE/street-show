/*
	dbcrud contains all code that interacts with a database;
	also holds the datastructures that the fetch api (handler) utilizes
*/
package dbcrud

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
)

// DB_info contains fields that correspond to the database configuration file, db_info.yaml
type DB_info struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     uint32 `yaml:"port"`
}

// User contains all fields that correspond to the User MySQL columns
type User struct {
	Username string `gorm"column:username" json:"username"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Id       uint64 `gorm:"column:id" json:"id"`
}

// Busker contains all fields that correspond to the Busker MySQL columns
type Busker struct {
	Username string `gorm"column:username" json:"username"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Id       uint64 `gorm:"column:id" json:"id"`
}

// DB is the instance of a gorm database
var DB *gorm.DB

// Open reads from the database config file (db_info.yaml), then accordingly establishes a localhost connection to the database
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

// SelectAll selects all busker rows from the Busker table
func (b *Busker) SelectAll() error {
	if err := DB.Find(&b).Error; err != nil {
		return err
	}
	return nil
}

// SelectLike selects all busker rows from the Busker table that share similar usernames or names to the likeString arg
func (b *Busker) SelectLike(likeString string) {
	DB.Where("username LIKE ? OR name LIKE ?", likeString, likeString).Find(&b)
}

// Create creates a new SQL Busker row
func (b *Busker) Create() {
	DB.Create(&b)
}

// SelectAll selects all user rows from the User table
func (u *User) SelectAll() error {
	if err := DB.Find(&u).Error; err != nil {
		return err
	}
	return nil
}

// SelectLike selects all user rows from the User table that share similar usernames or names to the likeString arg
func (u *User) SelectLike(likeString string) {
	DB.Where("username LIKE ? OR name LIKE ?", likeString, likeString).Find(&u)
}

// Create creates a new SQL User row
func (u *User) Create() {
	DB.Create(&u)
}

// ShowById selects the row of a user by searching for rows that share the index
func (u *User) ShowById(id int) {
	DB.Where("id=?", id).Find(&u)
}

// ShowById selects the row of a busker by searching for rows that share the index
func (b *Busker) ShowById(id int) {
	DB.Where("id=?", id).Find(&b)
}

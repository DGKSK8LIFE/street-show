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

// DB is the instance of a gorm database
var DB *gorm.DB

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
	Username string `gorm:"column:username" json:"username"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Id       uint64 `gorm:"column:id" json:"id"`
}

// Performance contains all fields that correspond  to the Performance MySQL columns
type Performance struct {
	Username string `gorm:"column:username" json:"username"`
	// will work out how I'll manage coord field later...
	Id uint64 `gorm:"column:id" json:"id"`
}

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
func (b *Busker) SelectLike(likeString string) error {
	if err := DB.Where("username LIKE ? OR name LIKE ?", likeString, likeString).Find(&b).Error; err != nil {
		return err
	}
	return nil
}

// Create creates a new SQL Busker row
func (b *Busker) Create() error {
	if err := DB.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

// ShowById selects the row of a busker by searching for rows that share the id
func (b *Busker) ShowById(id int) error {
	if err := DB.Where("id=?", id).Find(&b).Error; err != nil {
		return err
	}
	return nil
}

// SelectAll selects all user rows from the User table
func (u *User) SelectAll() error {
	if err := DB.Find(&u).Error; err != nil {
		return err
	}
	return nil
}

// SelectLike selects all user rows from the User table that share similar usernames or names to the likeString arg
func (u *User) SelectLike(likeString string) error {
	if err := DB.Where("username LIKE ? OR name LIKE ?", likeString, likeString).Find(&u).Error; err != nil {
		return err
	}
	return nil
}

// Create creates a new SQL User row
func (u *User) Create() error {
	if err := DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

// ShowById selects the row of a user by searching for rows that share the id
func (u *User) ShowById(id int) error {
	if err := DB.Where("id=?", id).Find(&u).Error; err != nil {
		return err
	}
	return nil
}

// SelectAll selects all performance rows from the Performance table
func (p *Performance) SelectAll() error {
	if err := DB.Find(&p).Error; err != nil {
		return err
	}
	return nil
}

// SelectLike selects all performance rows from the Performance table that share similar usernames to the likeString arg
func (p *Performance) SelectLike(likeString string) error {
	if err := DB.Where("username LIKE ?", likeString).Find(&p).Error; err != nil {
		return err
	}
	return nil
}

// Create creates a new SQL Performance row
func (p *Performance) Create() error {
	if err := DB.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

// ShowById selects the row of a Performance by searching for rows that share the id
func (p *Performance) ShowById(id int) error {
	if err := DB.Where("id=?", id).Find(&p).Error; err != nil {
		return err
	}
	return nil
}

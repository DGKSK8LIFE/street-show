package handler

import "github.com/gin-gonic/gin"

type User struct {
	Username string
	Name     string
	Email    string
	Id       uint64
}

func UserApi(c *gin.Context) {

}

package handler

import "github.com/gin-gonic/gin"

type UserAndBusker struct {
	Username string
	Name     string
	Email    string
	Id       uint64
}

func BuskerApi(c *gin.Context) {

}

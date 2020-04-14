package handler

import (
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
)

type UserAndBusker struct {
	Username string
	Name     string
	Email    string
	Id       uint64
}

func BuskerApi(c *gin.Context) {
	busker, err := url.QueryUnescape(c.DefaultQuery("busker", ""))
	if err != nil {
		log.Fatalf("QueryUnescape error: %s\n", err)
	}
}

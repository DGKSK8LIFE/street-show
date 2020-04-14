package handler

import (
	"log"
	"net/url"
	"street-show/app/backend/dbcrud"

	"github.com/gin-gonic/gin"
)

func BuskerApi(c *gin.Context) {
	busker, err := url.QueryUnescape(c.DefaultQuery("busker", ""))
	if err != nil {
		log.Fatalf("QueryUnescape error: %s\n", err)
	}
	scanTo := &dbcrud.Busker{}
	scanTo.SelectLike(busker)
	c.JSON(200, scanTo)
}

func UserApi(c *gin.Context) {
	user, err := url.QueryUnescape(c.DefaultQuery("user", ""))
	if err != nil {
		log.Fatalf("QueryUnescape error: %s\n", err)
	}
	scanTo := &dbcrud.User{}
	scanTo.SelectLikeUser(user)
	c.JSON(200, scanTo)
}

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
	dbcrud.SelectLike(scanTo, busker)
}

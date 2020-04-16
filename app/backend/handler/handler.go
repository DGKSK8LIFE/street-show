/*
	handler holds all of the routed functions directly called in main.go;
	handler is a fetch api which the React/React-Native view uses
*/
package handler

import (
	"log"
	"net/url"
	"street-show/app/backend/dbcrud"

	"github.com/gin-gonic/gin"
)

// BuskerApi serves json of buskers' data, utilizing the dbcrud module (mine)
func BuskerSearchApi(c *gin.Context) {
	busker, err := url.QueryUnescape(c.DefaultQuery("busker", ""))
	if err != nil {
		log.Fatalf("QueryUnescape error: %s\n", err)
	}
	scanTo := &dbcrud.Busker{}
	if len(busker) > 0 {
		scanTo.SelectLike(busker)
	}
	scanTo.SelectAll()
	c.JSON(200, scanTo)
}

// UserApi serves json of users' data, utilizing the dbcrud module (mine)
func UserSearchApi(c *gin.Context) {
	user, err := url.QueryUnescape(c.DefaultQuery("user", ""))
	if err != nil {
		log.Fatalf("QueryUnescape error: %s\n", err)
	}
	scanTo := &dbcrud.User{}
	if len(user) > 0 {
		scanTo.SelectLike(user)
	}
	scanTo.SelectAll()
	c.JSON(200, scanTo)
}

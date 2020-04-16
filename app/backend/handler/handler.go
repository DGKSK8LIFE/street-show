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
		c.JSON(200, scanTo)
		return
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
		c.JSON(200, scanTo)
		return
	}
	scanTo.SelectAll()
	c.JSON(200, scanTo)
}

// BuskerApi returns serialization of database row with precise id sent in a querystring
func BuskerApi(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	scanTo := &dbcrud.Busker{}
	if len(id) > 0 {
		scanTo.ShowById(id)
		c.JSON(200, scanTo)
		return
	}
	c.Status(404)
}

// UserApi returns serialization of database row with precise id sent in a querystring
func UserApi(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	scanTo := &dbcrud.User{}
	if len(id) > 0 {
		scanTo.ShowById(id)
		c.JSON(200, scanTo)
		return
	}
	c.Status(404)
}

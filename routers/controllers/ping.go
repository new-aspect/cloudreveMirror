package controllers

import (
	"cloudreveMirror/serializer"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Data: "Pong",
	})
}

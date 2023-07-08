package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	Shorten(c *gin.Context)
	Resolve(c *gin.Context)
}

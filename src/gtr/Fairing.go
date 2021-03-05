package gtr

import "github.com/gin-gonic/gin"

type Fairing interface {
	OnRequest(context *gin.Context) error
}
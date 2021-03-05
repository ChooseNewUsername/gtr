package middlewares

import (
	"github.com/gin-gonic/gin"
)

type UserMid struct {

}

func NewUserMid() *UserMid {
	return &UserMid{}
}

func (this *UserMid)OnRequest(context *gin.Context) error {

	return nil
}
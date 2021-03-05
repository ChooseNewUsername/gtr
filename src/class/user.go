package class

import (
	"github.com/gin-gonic/gin"
	"gtr/src/gtr"
)

type UserClass struct {

}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass)GetUser(ctx *gin.Context) string  {
	return "cc"
}
func (this *UserClass) Build(gtr *gtr.Gtr)  {
	gtr.Handle("GET","/user",this.GetUser)
}

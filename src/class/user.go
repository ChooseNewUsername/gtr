package class

import (
	"github.com/gin-gonic/gin"
	"gtr/src/gtr"
	"gtr/src/models"
)

type UserClass struct {

}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass)GetUser(ctx *gin.Context) string  {
	return "cc"
}
func (this *UserClass)GetUserLists(ctx *gin.Context) gtr.Model  {
	return &models.UserModel{UserId: 1,UserName: "ss"}
}
func (this *UserClass) Build(gtr *gtr.Gtr)  {
	gtr.Handle("GET","/user",this.GetUser)
	gtr.Handle("GET","/user2",this.GetUserLists)
}

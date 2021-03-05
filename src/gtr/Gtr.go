package gtr

import (
	"github.com/gin-gonic/gin"

)

type Gtr struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Ignite() *Gtr {
	return &Gtr{Engine:gin.New()}
}
func (this *Gtr) Launch()  {
	this.Run(":80")
}

func (this *Gtr) Handle(httpMethod string, relativePath string, handlers interface{}) *Gtr {
	//if h,ok:= handlers.(func(ctx *gin.Context) string);ok{
	//	this.g.Handle(httpMethod,relativePath, func(context *gin.Context) {
	//		context.JSON(200,gin.H{"data":h(context)})
	//	})
	//}
	if h:=Convert(handlers);h!=nil {
		this.g.Handle(httpMethod,relativePath,h)
	}
	return this
}

func (this *Gtr)Mount(group string,classes ...IClass) *Gtr  {
	this.g = this.Group(group)
	for _,class :=range classes{
		class.Build(this)
	}
	return this
}
func (this *Gtr)Attach(f Fairing) *Gtr  {
	this.Use(func(context *gin.Context) {
		err:=f.OnRequest(context)
		if err!=nil {
			context.AbortWithStatusJSON(400,gin.H{"err":err.Error()})
		}else {
			context.Next()
		}
	})
	return this
}
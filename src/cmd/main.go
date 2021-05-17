package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gtr/src/class"
	"gtr/src/gtr"
)

func main()  {


	gtr.Ignite().
		//Attach(NewUserMid()).
		Mount("cc",
			class.NewUserClass(),
			class.NewIndexClass()).
		Launch()

	g:=gin.New()
	g.Use(func(context *gin.Context) {
		fmt.Print(11)
	})
	g.Group("c")

	g.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"res":"ok",
		})
	})

	g.Run(":80")

}

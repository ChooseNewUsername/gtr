package class

import (
	"github.com/gin-gonic/gin"
	"gtr/src/gtr"
)

type IndexClass struct {

}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (this *IndexClass)GetIndex(ctx *gin.Context) string  {
	return "dd"
}
func (this *IndexClass) Build(gtr *gtr.Gtr)  {
	gtr.Handle("GET","/index",this.GetIndex)
	gtr.Handle("GET","/index2",this.GetIndex)
}
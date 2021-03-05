package gtr

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var ResponderList []Responder

func init()  {
	ResponderList = []Responder{
		new(StringResponder),
	}
}

type Responder interface {
	ResponderTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	h_ref:= reflect.ValueOf(handler)
	for _,r:=range ResponderList{
		r_ref := reflect.ValueOf(r).Elem()
		if h_ref.Type().ConvertibleTo(r_ref.Type()){
			r_ref.Set(h_ref)
			return r_ref.Interface().(Responder).ResponderTo()
		}
	}
	return nil
}

type StringResponder func( *gin.Context) string

func (this StringResponder)ResponderTo() gin.HandlerFunc  {
	return func(context *gin.Context) {
		context.String(200,this(context))
	}
}
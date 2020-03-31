package services

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(errCode int, msg string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func (g *Gin) StatusInternalServerError() {
	g.Ctx.JSON(500, Response{})
	return
}

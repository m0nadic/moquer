package util

import (
	"github.com/gin-gonic/gin"
)

type Payload struct {
	Ctx *gin.Context
}

func (p Payload) PP(key string) string {
	return p.Ctx.Param(key)
}
func (p Payload) PPQ(key string) string {
	return Quotify(p.Ctx.Param(key))
}





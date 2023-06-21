package main

import (
	"log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	pprof.Register(e)
	e.GET("/go", func(ctx *gin.Context) {
		//go func() {
		//
		//}()
		ctx.JSON(200, gin.H{
			"msg": "ok",
		})
	})

	log.Fatal(e.Run(":8888"))
}

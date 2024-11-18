package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("jenkins-test")
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		res := map[string]string{}
		res["data"] = "jenkins-test"
		ctx.JSON(http.StatusOK, res)
	})

	r.GET("/test", func(ctx *gin.Context) {
		res := map[string]string{}
		res["data"] = "test"
		ctx.JSON(http.StatusOK, res)
	})

	r.Run(":8080")
}

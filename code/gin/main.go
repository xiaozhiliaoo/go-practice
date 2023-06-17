package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozhiliaoo/go-practice/code/gin/middleware"
	"github.com/xiaozhiliaoo/go-practice/code/gin/service"
)

func main() {
	r := gin.New()
	r.POST("/hello",
		middleware.PreProcessing,
		service.Hello,
		middleware.PostProcessing)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

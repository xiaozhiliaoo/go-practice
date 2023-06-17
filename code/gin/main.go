package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/hello",
		PreProcessing,
		Hello,
		PostProcessing)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

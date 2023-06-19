package main

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	r := gin.New()
	r.POST("/hello",
		PreProcessing,
		Hello,
		PostProcessing)

	r.POST("/hello-timeout2",
		PreProcessing,
		timeout.New(
			timeout.WithTimeout(5*time.Second),
			timeout.WithHandler(Hello),
			timeout.WithResponse(PostProcessing),
		))

	r.POST("/hello-timeout", timeout.New(
		timeout.WithTimeout(100*time.Microsecond),
		timeout.WithHandler(emptySuccessResponse),
		timeout.WithHandler(nextFunc),
	))

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func emptySuccessResponse(c *gin.Context) {
	time.Sleep(30000 * time.Microsecond)
	c.String(http.StatusOK, "")
}

func nextFunc(c *gin.Context) {
	c.String(http.StatusOK, "quick timeout")
}

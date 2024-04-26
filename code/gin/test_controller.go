package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	i64 := c.Query("id")
	c.JSON(http.StatusOK, i64)
	return
}

func File(c *gin.Context) {

	response, err := http.Get("http://30.186.89.33/llm/evaluate_template.xlsx")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="evaluate_template.xlsx"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func GetParams(c *gin.Context) {
	i64 := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, i64+name)
	return
}

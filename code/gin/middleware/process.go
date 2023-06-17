package middleware

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

func PreProcessing(c *gin.Context) {

	id := c.Query("id")
	name := c.Query("name")

	log.Printf("id:%s,name:%s", id, name)

	c.Request.Body = ioutil.NopCloser(strings.NewReader(`{"question":"CVM和Lighthouse", "session_id":b9be4ddd7b17675e8eebc4b1a3de14db,"model_id":"51"}`))

	c.Set("is_hello", true)

}

// PostProcessing 响应后置处理
func PostProcessing(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	log.Printf("id:%s,name:%s", id, name)

	c.Writer.WriteString("测试流程通不通")
	return
}

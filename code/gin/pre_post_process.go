package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func PreProcessing(c *gin.Context) {

	id := c.Query("id")
	name := c.Query("name")

	log.Printf("id:%s,name:%s", id, name)

	req := HelloReq{ID: 1, Question: "123455"}
	reqByte, _ := json.Marshal(req)

	//c.Request.Body = ioutil.NopCloser(strings.NewReader(`{"question":"CVM和Lighthouse", "session_id":b9be4ddd7b17675e8eebc4b1a3de14db,"model_id":"51"}`))
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqByte))

	c.Set("is_hello", true)

}

// PostProcessing 响应后置处理
func PostProcessing(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	log.Printf("id:%s,name:%s", id, name)

	answer := c.Value("hello_answer").(HelloReq)

	answerByte, _ := json.Marshal(answer)

	c.Writer.Write(answerByte)

	return
}

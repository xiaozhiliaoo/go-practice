package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

func PreProcessing(c *gin.Context) {
	log.Print("Entry PreProcessing")

	id := c.Query("id")
	name := c.Query("name")

	log.Printf("PreProcessing id:%s,name:%s", id, name)

	req := HelloReq{ID: 1, Question: "123455"}
	reqByte, _ := json.Marshal(req)

	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqByte))

	c.Set("is_hello", true)

}

// PostProcessing 响应后置处理
func PostProcessing(c *gin.Context) {

	log.Printf("timeout go to")

	log.Printf("PostProcessing answer11111：%s", c.GetString("111"))

	answer1 := c.GetString("hello_answer22")
	log.Printf("PostProcessing answer1：%s", answer1)

	id := c.Query("id")
	name := c.Query("name")
	log.Printf("PostProcessing id:%s,name:%s", id, name)

	//answer := c.Value("hello_answer").(HelloReq)

	//answerByte, _ := json.Marshal(answer)

	c.Writer.WriteString("ok result")

	return
}

// PostProcessingTimeOut 响应后置处理
func PostProcessingTimeOut(c *gin.Context) {
	time.Sleep(10 * time.Second)

	log.Printf("timeout go to")

	id := c.Query("id")
	name := c.Query("name")
	log.Printf("PostProcessingTimeOut id:%s,name:%s", id, name)

	c.Writer.WriteString("timeout result")

	return
}

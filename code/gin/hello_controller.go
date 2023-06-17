package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

type HelloReq struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Date     string `json:"date"`
}

func Hello(c *gin.Context) {
	isHello := c.Value("is_hello").(bool)
	log.Printf("isHello:%t", isHello)

	req := HelloReq{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		//body, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("ShouldBind error:%+v", err)
		c.Writer.WriteString("ShouldBind失败")
		return
	}

	log.Printf("hello body:%+v", req)

	c.Set("hello_answer", req)

	c.Writer.WriteString(c.Query("id") + "---" + c.Query("name"))

}

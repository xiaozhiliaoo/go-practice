package service

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type HelloReq struct {
	ModelID   int    `json:"model_id"`
	SessionID string `json:"session_id"`
	Question  string `json:"question"`
}

func Hello(c *gin.Context) {
	isHello := c.Value("is_hello").(bool)
	log.Printf("isHello:%t", isHello)

	req := HelloReq{}
	if err := c.ShouldBind(&req); err != nil {
		body, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("ShouldBind error:%+v,body:%+v", err, body)
		c.Writer.WriteString("ShouldBind失败")
		return
	}

	log.Printf("hello body:%+v", req)

	c.Writer.WriteString(c.Query("id") + "---" + c.Query("name"))

}

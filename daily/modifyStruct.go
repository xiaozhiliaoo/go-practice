package main

import "fmt"

type Req struct {
	Answer   string `json:"answer"`
	RecordID int64  `json:"record_id"`
}

type WrapperReq struct {
	Request *Req
}

func main() {
	//req := Req{}
	//modify(&req)
	req := []Req{{"111", 11}, {"222", 222}}
	modifySlice(req)
	fmt.Printf("%+v", req)
}

func modifySlice(req []Req) {
	for i := range req {
		req[i].Answer += "jiandan"
	}
}

func modify(req *Req) {
	req.Answer = "111"
	req.RecordID = 1111
	modify2(&WrapperReq{
		Request: req,
	})
}

func modify2(req *WrapperReq) {
	req.Request.Answer = "22ss22"
	req.Request.RecordID = 2222
}

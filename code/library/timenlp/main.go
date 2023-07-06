package main

import (
	"fmt"
	timenlp "github.com/bububa/TimeNLP"
	utils "github.com/xiaozhiliaoo/common-golang/common-utils"
	"time"
)

func main() {
	now := time.Now()
	r, err := timenlp.NewTimeNormalizer(false).Parse("这个月最后一天", now)
	if err != nil {
		fmt.Printf("error:%+v", err)
	}
	if len(r.Points) > 0 {
		ti := r.Points[0].Time
		str := utils.ConvertTime2Str(ti)
		fmt.Printf("%s", str)
	}
}

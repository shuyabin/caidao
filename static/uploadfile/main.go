package main

import (
	"log"
	"time"

	"github.com/imroc/req"
)

func main1() {
	url := "http://www.jrjingshan.com/"
	req.SetFlags(req.LstdFlags | req.Lcost) // 输出格式显示请求耗时
	r, _ := req.Get(url)
	log.Println(r)                // http://foo.bar/api 3.260802ms {"code":0 "msg":"success"}
	if r.Cost() > 3*time.Second { // 检查耗时
		log.Println("WARN: slow request:", r)
	}
	// fmt.Print(res.Cost())
	// resp := res.Response()
	// fmt.Println(resp)

}

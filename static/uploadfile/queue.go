package main

import (
	"fmt"
	"sync"

	"github.com/Damnever/goqueue"
	"github.com/PuerkitoBio/goquery"
	"github.com/henrylee2cn/pholcus/common/mahonia"
	"github.com/imroc/req"
)

func main() {
	// queue := goqueue.New(0)
	// wg := &sync.WaitGroup{}

	// for i := 0; i <= 1000000; i++ {
	// 	queue.PutNoWait("http://www.jrjingshan.com/id=" + strconv.Itoa(i))
	// }

	addURL("http://www.jrjingshan.com/")
	// addUrl 添加url到队列  && 判断queue是否存在这个url 当前不存在的链接才添加
	// 返回的数据有 : 扫描总数链接多少, 外链链接多少 , 错误链接多少, 超时链接多少
	// 错误链接的 标题. 连接
	// start := time.Now()
	// for i := 0; i < 5; i++ {
	// 	go worker(queue, wg)
	// }
	// wg.Add(5)
	// wg.Wait()
	// end := time.Now()

	// fmt.Println(start)
	// fmt.Println(end)
	// fmt.Println("All task done!!!")
}

//addURL 添加url
func addURL(domen string) {
	url := "http://www.jrjingshan.com/"
	r, _ := req.Get(url)
	dec := mahonia.NewDecoder("UTF-8")
	rd := dec.NewReader(r.Response().Body)
	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		fmt.Print(err)
	}
	urlArr := doc.Find("a")
	urlArr.Each(func(i int, content *goquery.Selection) {
		// a, _ := content.Attr("href")
		title := content.Text()
		fmt.Println(title)
	})
}

func worker(queue *goqueue.Queue, wg *sync.WaitGroup) {
	defer wg.Done()
	for !queue.IsEmpty() {
		val, err := queue.Get(0)
		if err != nil {
			fmt.Println("Unexpect Error: %v\n", err)
		}
		fmt.Printf("-> %v\n", val)
	}
}

package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/imroc/req"
)

func WriteAll(filePth string, data string) {
	d1 := []byte(data)
	err := ioutil.WriteFile(filePth, d1, 0644)
	if err != nil {
		fmt.Println("文件写入失败")
	}
}
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

// 启动post请求发送给 目标server端进行交互数据
func Startpost(url string, pwd string, action string, pamone string, pamtwo string) string {
	//这里进行post发送数据和接收返回的数据
	param := req.Param{
		"action": action,
		pwd:      "1",
		"z1":     pamone,
		"z2":     pamtwo,
	}
	//	form-data; name="upload"; filename="timg.jpg"

	r, err := req.Post(url, param)

	if err != nil {
		fmt.Println(err)
	}
	resp := r.String()
	return resp
}

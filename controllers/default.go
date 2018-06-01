package controllers

import (
	"caidao/models"
	"caidao/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tidwall/gjson"
)

type MainController struct {
	beego.Controller
}

// 根据url中的参数获取密码、和URL
func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

var stu *models.Caidao

func init() {
	stu = new(models.Caidao)

}

//Listjson  url列表
func (c *MainController) Listjson() {
	o := orm.NewOrm()
	var caidao []*models.Caidao
	o.Raw("SELECT *  FROM Caidao").QueryRows(&caidao)
	caidao_json, _ := json.Marshal(caidao)
	c.Ctx.WriteString(string(caidao_json))
}

// List 模板显示
func (c *MainController) Wenjian() {
	c.TplName = "wenjian.tpl"
}

//Add 处理添加
func (c *MainController) Add() {
	o := orm.NewOrm()
	stu.URL = c.GetString("url")
	stu.Pwd = c.GetString("pwd")
	stu.Seting = c.GetString("seting")
	o.Insert(stu)
	// 皮一下..就不判断了
	c.Ctx.WriteString("{'error': \"0\",\"msg\": \"添加成功\"}")
}

//Delete 删除
func (c *MainController) DeleteUrl() {
	id := c.GetString("id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	result := stu.DeleteByid(intid)
	c.Ctx.WriteString(result)
}

// Addurl 添加URL
func (c *MainController) Addurl() {
	c.TplName = "add.tpl"
}

// Filelist 文件列表
func (c *MainController) Filelist() {
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	result := stu.FindOne(intid)
	fileinfo := util.Startpost(result.URL, result.Pwd, "readdict", " ", " ")
	c.Ctx.WriteString(fileinfo)
}

// 搜索文件
func (c *MainController) Search() {
	source := c.GetString("source")
	target := c.GetString("text")
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	result := stu.FindOne(intid)
	fileinfo := util.Startpost(result.URL, result.Pwd, "search", source, target)
	c.Ctx.WriteString(fileinfo)
}

// 移动文件
func (c *MainController) Move() {
	source := c.GetString("source")
	target := c.GetString("target")
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	result := stu.FindOne(intid)
	fileinfo := util.Startpost(result.URL, result.Pwd, "move", source, target)
	c.Ctx.WriteString(fileinfo)
}

// 复制文件
func (c *MainController) Copy() {
	source := c.GetString("source")
	target := c.GetString("target")
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	result := stu.FindOne(intid)
	fileinfo := util.Startpost(result.URL, result.Pwd, "copy", source, target)
	c.Ctx.WriteString(fileinfo)
}

// 重命名
func (c *MainController) Rename() {
	source := c.GetString("source")
	target := c.GetString("target")
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	result := stu.FindOne(intid)
	res := util.Startpost(result.URL, result.Pwd, "rename", source, target)
	c.Ctx.WriteString(res)
}

// 下载
func (c *MainController) Download() {
	source := c.GetString("source")
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	res := stu.FindOne(intid)
	result := util.Startpost(res.URL, res.Pwd, "download", source, "")
	//下载到go服务器上再下载到本地
	fileName := gjson.Get(result, "filename")
	content := gjson.Get(result, "content")
	decodeBytes, _ := base64.StdEncoding.DecodeString(content.String())
	util.WriteAll("static/download/"+fileName.String(), string(decodeBytes))
	//下载文件
	c.Ctx.Output.Download("static/download/" + fileName.String())
}

//删除
func (c *MainController) Delete() {
	source := c.GetString("source")
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	res := stu.FindOne(intid)
	result := util.Startpost(res.URL, res.Pwd, "delete", source, "")
	c.Ctx.WriteString(result)
}

// 创建文件夹
func (c *MainController) NewFolder() {
	source := c.GetString("source")
	target := c.GetString("target")
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	res := stu.FindOne(intid)

	result := util.Startpost(res.URL, res.Pwd, "create", source, target)
	c.Ctx.WriteString(result)
}

// 上传
func (c *MainController) Fileupload() {
	target := c.GetString("target")
	f, h, _ := c.GetFile("upload")
	fileName := h.Filename
	//	arr := strings.Split(fileName, ":")
	id := c.GetString("d_id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("字符串转换错误")
	}
	res := stu.FindOne(intid)
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	err = c.SaveToFile("upload", path.Join("static/uploadfile", fileName))
	fmt.Println(err)
	if err == nil {
		fileinfo, _ := util.ReadAll("static/uploadfile/" + fileName)
		result := util.Startpost(res.URL, res.Pwd, "upload", target+"/"+fileName, string(fileinfo))
		c.Ctx.WriteString(result)
	} else {
		c.Ctx.WriteString("upload error")
	}
}

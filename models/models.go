package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// type Article struct {
// 	Id     int    `form:"-"`
// 	Name   string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
// 	Client string `form:"client,text,client:"`
// 	Url    string `form:"url,text,url:"`
// }
type Caidao struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Pwd    string `json:"pwd"`
	Seting string `json:"seting"`
}

//FindOne 通过id找到URL 和密码
func (a *Caidao) FindOne(id int) Caidao {
	o := orm.NewOrm()
	result := Caidao{ID: id}
	err := o.Read(&result)
	if err != nil {
		fmt.Print(err)
	}
	return result
}

// DeleteByid 删除
func (a *Caidao) DeleteByid(id int) string {
	o := orm.NewOrm()
	if _, err := o.Delete(&Caidao{ID: id}); err == nil {
		return "{'error': \"0\",\"msg\": \"删除成功\"}"
	}
	return "{'error': \"400\",\"msg\": \"删除失败\"}"
}

// TableName 返回表名称
func (a *Caidao) TableName() string {
	return "Caidao"
}

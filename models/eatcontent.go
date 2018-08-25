package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
)

type EatContent struct {
	Id         int64  `orm:"auto;unique;column(id)"`
	Content    string `orm:"column(content);size(64)"`
	CreateTime int64
}

func (eat *EatContent) TableName() string {
	return "eat_content"
}

func NewEatContent() *EatContent {
	return &EatContent{}
}

// 创建
func (eat *EatContent) Insert(content string) (id int64 , err error) {
	content = strings.TrimSpace(content)
	eatModel , err := eat.FindContent(content)
	if err == nil  {
		return eatModel.Id , err
	}
	o := orm.NewOrm()
	eatModel.Content = content
	eatModel.CreateTime = time.Now().Unix()
	return o.Insert(eat)
}

// 查询
func (eat *EatContent) Find(id int64, cols ...string) (*EatContent, error) {
	if id <= 0 {
		return eat, InvalidId
	}
	o := orm.NewOrm()
	err := o.QueryTable(eat.TableName()).Filter("id", id).One(eat, cols...)
	return eat, err
}


// 查询
func (eat *EatContent) FindContent(content string, cols ...string) (*EatContent, error) {
	if content == "" {
		return eat, InvalidId
	}
	o := orm.NewOrm()
	err := o.QueryTable(eat.TableName()).Filter("content", content).One(eat, cols...)
	return eat, err
}

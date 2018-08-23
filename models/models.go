package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int64  `orm:"auto;unique;column(id)"`
	Openid     string `orm:"column(openid);size(64);unique"`
	Mobile     string `orm:"column(mobile);size(32)"`
	Email      string `orm:"column(email);size(32)"`
	Nickname   string `orm:"column(nickname);size(32)"`
	Avatar     string `orm:"column(avatar);size(32)"`
	CityText   string `orm:"column(city_text);size(64)"`
	Lat        string `orm:"column(lat);size(32)"`
	Lon        string `orm:"column(lon);size(32)"`
	Sex        uint
	CreateTime int64
}

func (user *User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}

// 创建用户
func (user *User) Insert() error {
	o := orm.NewOrm()
	user.CreateTime = time.Now().Unix()
	_, err := o.Insert(user)
	return err
}

// 查询用户
func (user *User) Find(id int64, cols ...string) (*User, error) {
	if id <= 0 {
		return user, InvalidId
	}
	o := orm.NewOrm()
	err := o.QueryTable(user.TableName()).Filter("id", id).One(user, cols...)
	return user, err
}

// 更新用户
func (user *User) Update(cols ...string) error {
	o := orm.NewOrm()
	temp := NewUser()
	temp.Id = user.Id
	if err := o.Read(temp); err != nil {
		return err
	}
	_, err := o.Update(user, cols...)
	return err
}

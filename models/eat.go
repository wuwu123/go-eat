package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"github.com/astaxie/beego"
)

type Eat struct {
	Id         int64 `orm:"auto;unique;column(id)"`
	UserID     int64 `orm:"column(user_id);size(64)"`
	ContentId  int64 `orm:"column(content_id)"`
	Num        uint
	Status     uint
	CreateTime int64
}

func (eat *Eat) TableName() string {
	return "user_eat"
}

func NewEat() *Eat {
	return &Eat{}
}

// 创建
func (eat *Eat) Insert() error {
	o := orm.NewOrm()
	eatModel, err := eat.FindByUserEat(eat.UserID, eat.ContentId , "id")
	if eatModel.Id == 0 {
		eat.CreateTime = time.Now().Unix()
		_, err = o.Insert(eat)
	}
	return err
}

// 查询
func (eat *Eat) Find(id int64, cols ...string) (*Eat, error) {
	if id <= 0 {
		return eat, InvalidId
	}
	o := orm.NewOrm()
	err := o.QueryTable(eat.TableName()).Filter("id", id).One(eat, cols...)
	return eat, err
}

// 更新
func (eat *Eat) Update(cols ...string) error {
	o := orm.NewOrm()
	temp := NewEat()
	temp.Id = eat.Id
	if err := o.Read(temp); err != nil {
		return err
	}
	_, err := o.Update(eat, cols...)
	return err
}

// 查询用户
func (eat *Eat) FindByUser(user_id int64, page int64, limit int64) (eats []*Eat, totalCount int, err error) {
	o := orm.NewOrm()

	offset := (page - 1) * limit
	// 原生sql查询
	sql := fmt.Sprintf("SELECT COUNT(*) FROM %s where user_id = ?", eat.TableName())
	err = o.Raw(sql, user_id).QueryRow(&totalCount)
	if err != nil {
		return
	}
	// 高级查询
	qs := o.QueryTable(eat.TableName())
	qs.Filter("user_id", user_id)
	qs.Limit(limit, offset)
	qs.OrderBy("-id")
	_, err = qs.RelatedSel().All(&eats)
	beego.Info("我是结果" , eats)
	return
}

// 查询 菜单
func (eat *Eat) FindByUserEat(userId int64, contentId int64, cols ...string) (*Eat, error) {
	o := orm.NewOrm()
	err := o.QueryTable(eat.TableName()).Filter("user_id", userId).Filter("content_id", contentId).One(eat, cols...)
	beego.Error(eat)
	return eat, err
}

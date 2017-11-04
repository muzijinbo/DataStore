package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Phone   string
	Created time.Time `orm:"index"`
	Level   int32     `orm:"index"`
	UType   int32
}

type Data struct {
	Id         int64
	Location   string
	Name1      string
	Name2      string
	Name3      string
	Primary    string
	Number     int
	Created    time.Time `orm:"index"` //创建时间
	Size       float32
	Introduce  string
	Attachment string `orm:"size(5000)"`
	DType      string
}

func RegisterDB() {
	// 需要在init中注册定义的model

	orm.RegisterModel(new(User), new(Data))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:111@/myapp?charset=utf8")
}

func GetUser() ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.All(&users)
	return users, err
}
func GetHeaderData() ([]*Data, error) {
	o := orm.NewOrm()
	datas := make([]*Data, 0)
	qs := o.QueryTable("data")
	_, err := qs.All(&datas)
	return datas, err
}

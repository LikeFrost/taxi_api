package models

import "github.com/astaxie/beego/orm"

type Money struct {
	Id      int    `orm:"pk;auto"`
	Type    string //in or out
	InTime  string
	InType  string
	InMoney string
	People  string
	More    string
}

func AddMoney(types, in_time, in_type, in_money, people, more string) (code int, msg string) {
	o := orm.NewOrm()
	m := Money{}
	err := o.Read(&m)
	if err == nil {
		return 103, "收支信息已存在"
	} else {
		m.Type = types
		m.InTime = in_time
		m.InType = in_type
		m.InMoney = in_money
		m.People = people
		m.More = more
		_, err = o.Insert(&m)
		if err == nil {
			return 100, "创建收支信息成功"
		} else {
			return 101, "创建收支信息失败"
		}
	}
}
func DeleteMoney(id int) (code int, msg string) {
	o := orm.NewOrm()
	m := Money{Id: id}
	err := o.Read(&m)
	if err == nil {
		_, err = o.Delete(&m)
		if err == nil {
			return 100, "删除收支信息成功"
		}
	}
	return 101, "删除收支信息失败"
}
func GetAllMoney() (code int, msg string, m []*Money) {
	o := orm.NewOrm()
	_, err := o.QueryTable("money").All(&m)
	if err == nil {
		return 100, "获取全部收支信息成功", m
	} else {
		return 101, "获取全部收支信息失败", m
	}
}

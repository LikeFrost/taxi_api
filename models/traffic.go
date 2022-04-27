package models

import "github.com/astaxie/beego/orm"

type Traffic struct {
	Id       int `orm:"pk;auto"`
	CarId    string
	User     string
	Content  string
	Time     string
	Location string
	Money    string
	Score    string
	State    string
}

func AddTraffic(car_id, user, content, time, location, money, score, state string) (code int, msg string) {
	o := orm.NewOrm()
	t := Traffic{}
	t.CarId = car_id
	t.User = user
	t.Content = content
	t.Time = time
	t.Location = location
	t.Money = money
	t.Score = score
	t.State = state
	_, err := o.Insert(&t)
	if err == nil {
		return 100, "添加违章信息成功"
	} else {
		return 101, "添加违章信息失败"
	}
}
func GetTrafficByCarId(car_id string) (code int, msg string, t []*Traffic) {
	o := orm.NewOrm()
	_, err := o.QueryTable("traffic").Filter("User__in", car_id).All(&t)
	if err == nil {
		return 100, "获取违章信息成功", t
	} else {
		return 101, "获取违章信息失败", t
	}
}
func DeleteTraffic(id int) (code int, msg string) {
	o := orm.NewOrm()
	t := Traffic{Id: id}
	err := o.Read(&t)
	if err == nil {
		_, err = o.Delete(&t)
		if err == nil {
			return 100, "删除违章信息成功"
		}
	}
	return 101, "删除违章信息失败"
}
func GetAllTraffic() (code int, msg string, t []*Traffic) {
	o := orm.NewOrm()
	_, err := o.QueryTable("traffic").All(&t)
	if err == nil {
		return 100, "获取全部违章信息成功", t
	} else {
		return 101, "获取全部违章信息失败", t
	}
}

func GetSearchTraffic(search string) (t []*Traffic, code int, msg string) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("Id__exact", search)
	qs := o.QueryTable("traffic")
	qs = qs.SetCond(cond1)
	qs.All(&t)
	return t, 100, "查询成功"
}

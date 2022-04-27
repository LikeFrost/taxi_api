package models

import "github.com/astaxie/beego/orm"

type Question struct {
	Id         int `orm:"pk;auto"`
	CarId      string
	HappenTime string
	CreateTime string
	Method     string
	Content    string
	Result     string
}

func AddQuestion(car_id, happen_time, create_time, method, content, result string) (code int, msg string) {
	o := orm.NewOrm()
	q := Question{}
	q.CarId = car_id
	q.HappenTime = happen_time
	q.CreateTime = create_time
	q.Method = method
	q.Content = content
	q.Result = result
	_, err := o.Insert(&q)
	if err == nil {
		return 100, "添加投诉信息成功"
	} else {
		return 101, "添加投诉信息失败"
	}
}
func GetQuestionByCarId(car_id string) (code int, msg string, q []*Question) {
	o := orm.NewOrm()
	_, err := o.QueryTable("question").Filter("CarId__in", car_id).All(&q)
	if err == nil {
		return 100, "获取投诉信息成功", q
	} else {
		return 101, "获取投诉信息失败", q
	}
}
func DeleteQuestion(id int) (code int, msg string) {
	o := orm.NewOrm()
	q := Question{Id: id}
	err := o.Read(&q)
	if err == nil {
		_, err = o.Delete(&q)
		if err == nil {
			return 100, "删除投诉信息成功"
		}
	}
	return 101, "删除投诉信息失败"
}
func GetAllQuestion() (code int, msg string, q []*Question) {
	o := orm.NewOrm()
	_, err := o.QueryTable("question").All(&q)
	if err == nil {
		return 100, "获取全部投诉信息成功", q
	} else {
		return 101, "获取全部投诉信息失败", q
	}
}

func GetSearchQuestion(search string) (q []*Question, code int, msg string) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("Id__exact", search).Or("CreateTime__exact", search)
	qs := o.QueryTable("question")
	qs = qs.SetCond(cond1)
	qs.All(&q)
	return q, 100, "查询成功"
}

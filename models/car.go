package models

import "github.com/astaxie/beego/orm"

type Car struct {
	CarId         string `orm:"pk"`
	Color         string
	CarType       string
	Id            string
	CountMoney    string
	State         string
	EnsureCompany string
	EnsureYear    string
	SafeDay       string
	EnsureDay     string
}

func AddCar(car_id, color, car_type, id, count_money, state, ensure_company, ensure_year, safe_day, ensure_day string) (code int, msg string) {
	o := orm.NewOrm()
	c := Car{CarId: car_id}
	err := o.Read(&c)
	if err == nil {
		return 103, "车辆信息已存在"
	} else {
		c.Color = color
		c.CarType = car_type
		c.Id = id
		c.CountMoney = count_money
		c.State = state
		c.EnsureCompany = ensure_company
		c.EnsureYear = ensure_year
		c.SafeDay = safe_day
		c.EnsureDay = ensure_day
		_, err = o.Insert(&c)
		if err == nil {
			return 100, "创建车辆信息成功"
		} else {
			return 101, "创建车辆信息失败"
		}
	}
}
func UpdateCar(car_id, color, car_type, id, count_money, state, ensure_company, ensure_year, safe_day, ensure_day string) (code int, msg string) {
	o := orm.NewOrm()
	c := Car{CarId: car_id}
	err := o.Read(&c)
	if err == nil {
		c.Color = getResult(color, c.Color)
		c.CarType = getResult(car_type, c.CarType)
		c.Id = getResult(id, c.Id)
		c.CountMoney = getResult(count_money, c.CountMoney)
		c.State = getResult(state, c.State)
		c.EnsureCompany = getResult(ensure_company, c.EnsureCompany)
		c.EnsureYear = getResult(ensure_year, c.EnsureYear)
		c.SafeDay = getResult(safe_day, c.SafeDay)
		c.EnsureDay = getResult(ensure_day, c.EnsureDay)
		_, err = o.Update(&c)
		if err == nil {
			return 100, "更新车辆信息成功"
		} else {
			return 101, "更新车辆信息失败"
		}
	}
	return 103, "车辆信息不存在"
}

func GetCar(car_id string) (code int, msg string, car Car) {
	o := orm.NewOrm()
	car = Car{CarId: car_id}
	err := o.Read(&car)
	if err == nil {
		return 100, "获取车辆信息成功", car
	} else {
		return 101, "获取车辆信息失败", car
	}
}
func DeleteCar(car_id string) (code int, msg string) {
	o := orm.NewOrm()
	car := Car{CarId: car_id}
	err := o.Read(&car)
	if err == nil {
		_, err = o.Delete(&car)
		if err == nil {
			return 100, "删除车辆信息成功"
		}
	}
	return 101, "删除车辆信息失败"
}
func GetAllCar() (code int, msg string, c []*Car) {
	o := orm.NewOrm()
	_, err := o.QueryTable("car").All(&c)
	if err == nil {
		return 100, "获取全部车辆信息成功", c
	} else {
		return 101, "获取全部车辆信息失败", c
	}
}

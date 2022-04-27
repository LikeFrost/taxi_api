package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Name      string `orm:"pk"`
	Id        string
	Password  string
	Auth      string
	Sex       string
	IdCard    string
	Telephone string
	Birthday  string
	Home      string
	CarId     string
	DriveId   string
	DriveCar  string
	DriveYear string
}

//注册
func LogUp(name, password, auth string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Name: name}
	err := o.Read(&u)
	if err == nil {
		return 101, "用户已存在,请直接登录!"
	} else {
		u.Password = password
		u.Auth = auth
		_, err = o.Insert(&u)
		if err == nil {
			return 100, "注册成功!"
		} else {
			return 102, "注册失败,请稍后再试!"
		}
	}
}

//登录
func Login(name, password string) (code int, msg, auth string) {
	o := orm.NewOrm()
	u := User{Name: name}
	err := o.Read(&u)
	if err != nil {
		return 101, "用户不存在,请注册后登录!", ""
	} else if u.Password != password {
		return 102, "密码错误,请检查后重试!", ""
	}
	return 100, "登录成功!", u.Auth
}

func getResult(x, y string) (result string) {
	if x == "" {
		return y
	} else {
		return x
	}
}

//更新用户信息
func UpdateUser(name, id, sex, id_card, telephone, birthday, home, car_id, drive_id, drive_car, drive_year string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Name: name}
	err := o.Read(&u)
	if err == nil {
		u.Id = getResult(id, u.Id)
		u.Sex = getResult(sex, u.Sex)
		u.IdCard = getResult(id_card, u.IdCard)
		u.Telephone = getResult(telephone, u.Telephone)
		u.Birthday = getResult(birthday, u.Birthday)
		u.Home = getResult(home, u.Home)
		u.CarId = getResult(car_id, u.CarId)
		u.DriveId = getResult(drive_id, u.DriveId)
		u.DriveCar = getResult(drive_car, u.DriveCar)
		u.DriveYear = getResult(drive_year, u.DriveYear)
		_, err := o.Update(&u)
		if err == nil {
			return 100, "更新信息成功"
		} else {
			return 101, "更新信息失败"
		}
	}
	return 105, "用户不存在"
}

//更新用户信息
func UpdateOtherUser(name, id, auth, sex, id_card, telephone, birthday, home, car_id, drive_id, drive_car, drive_year string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Name: name}
	err := o.Read(&u)
	if err == nil {
		u.Id = getResult(id, u.Id)
		u.Auth = getResult(auth, u.Auth)
		u.Sex = getResult(sex, u.Sex)
		u.IdCard = getResult(id_card, u.IdCard)
		u.Telephone = getResult(telephone, u.Telephone)
		u.Birthday = getResult(birthday, u.Birthday)
		u.Home = getResult(home, u.Home)
		u.CarId = getResult(car_id, u.CarId)
		u.DriveId = getResult(drive_id, u.DriveId)
		u.DriveCar = getResult(drive_car, u.DriveCar)
		u.DriveYear = getResult(drive_year, u.DriveYear)
		_, err := o.Update(&u)
		if err == nil {
			return 100, "更新信息成功"
		} else {
			return 101, "更新信息失败"
		}
	}
	return 105, "用户不存在"
}

//获取用户信息
func GetUser(name string) (User, bool) {
	o := orm.NewOrm()
	u := User{Name: name}
	err := o.Read(&u)
	if err == nil {
		return u, true
	}
	return User{}, false
}
func DeleteUser(name string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Name: name}
	err := o.Read(&u)
	if err == nil {
		_, e := o.Delete(&u)
		if e == nil {
			return 100, "删除成功"
		}
	}

	return 101, "删除失败"
}

func GetAllUsers() (u []*User, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("user").Exclude("Auth__in", "super").All(&u)
	if err == nil {
		return u, 100, "获取全部用户成功"
	} else {
		return u, 101, "获取全部用户失败"
	}
}

package controllers

import (
	"encoding/json"
	"taxi_api/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description 登录
// @router /login [post]
func (u *UserController) Login() {
	var data map[string]interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &data)
	name := data["name"].(string)
	password := data["password"].(string)
	code, msg, auth := models.Login(name, password)
	if code == 100 {
		//创建token
		claims := make(jwt.MapClaims)
		claims["name"] = name
		claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("password"))
		if err == nil {
			u.Data["json"] = map[string]interface{}{
				"code":  code,
				"msg":   msg,
				"auth":  auth,
				"token": tokenString,
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code":  103,
				"msg":   msg,
				"auth":  auth,
				"token": "生成失败",
			}
		}
	} else {
		u.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
		}
	}
	u.ServeJSON()
}

// @Title LogUp
// @Description 注册
// @router /logup [post]
func (u *UserController) LogUp() {
	var data map[string]interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &data)
	name := data["name"].(string)
	password := data["password"].(string)
	auth := data["auth"].(string)
	code, msg := models.LogUp(name, password, auth)
	if code == 100 {
		//创建token
		claims := make(jwt.MapClaims)
		claims["name"] = name
		claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("password"))
		if err == nil {
			u.Data["json"] = map[string]interface{}{
				"code":  code,
				"msg":   msg,
				"token": tokenString,
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code":  103,
				"msg":   msg,
				"token": "生成失败",
			}
		}
	} else {
		u.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
		}
	}
	u.ServeJSON()
}

// @Title UpdateUser
// @Description 更新用户信息
// @router / [post]
func (u *UserController) UpdateUser() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		var data map[string]interface{}
		json.Unmarshal(u.Ctx.Input.RequestBody, &data)
		id := data["id"].(string)
		sex := data["sex"].(string)
		id_card := data["id_card"].(string)
		telephone := data["telephone"].(string)
		birthday := data["birthday"].(string)
		home := data["home"].(string)
		car_id := data["car_id"].(string)
		drive_id := data["drive_id"].(string)
		drive_car := data["drive_car"].(string)
		drive_year := data["drive_year"].(string)
		if !ok {
			u.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.UpdateUser(name, id, sex, id_card, telephone, birthday, home, car_id, drive_id, drive_car, drive_year)
			u.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	u.ServeJSON()
}

// @Title UpdateUserByName
// @Description 更新其他用户信息
// @router /updateother [post]
func (u *UserController) UpdateUserByName() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		var data map[string]interface{}
		json.Unmarshal(u.Ctx.Input.RequestBody, &data)
		name := data["name"].(string)
		id := data["id"].(string)
		auth := data["auth"].(string)
		sex := data["sex"].(string)
		id_card := data["id_card"].(string)
		telephone := data["telephone"].(string)
		birthday := data["birthday"].(string)
		home := data["home"].(string)
		car_id := data["car_id"].(string)
		drive_id := data["drive_id"].(string)
		drive_car := data["drive_car"].(string)
		drive_year := data["drive_year"].(string)
		if !ok {
			u.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.UpdateOtherUser(name, id, auth, sex, id_card, telephone, birthday, home, car_id, drive_id, drive_car, drive_year)
			u.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description 获取用户自身信息
// @router / [get]
func (u *UserController) Get() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		u.ServeJSON()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	user, flag := models.GetUser(name)
	if !ok {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		if flag {
			u.Data["json"] = map[string]interface{}{
				"code": 100,
				"msg":  "获取信息成功",
				"user": user,
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code": 101,
				"msg":  "获取信息失败",
			}
		}
	}
	u.ServeJSON()
}

// @Title GetUserByName
// @Description 获取其他用户信息
// @router /:name [get]
func (u *UserController) GetUserByName() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		u.ServeJSON()
		return
	}
	_, ok := token.Claims.(jwt.MapClaims)
	name := u.GetString(":name")
	user, flag := models.GetUser(name)
	if !ok {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		if flag {
			u.Data["json"] = map[string]interface{}{
				"code": 100,
				"msg":  "获取信息成功",
				"user": user,
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code": 101,
				"msg":  "获取信息失败",
			}
		}
	}
	u.ServeJSON()
}

// @Title DeleteUser
// @Description 删除用户
// @router /:name [Delete]
func (u *UserController) DeleteUser() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		u.ServeJSON()
		return
	}
	_, ok := token.Claims.(jwt.MapClaims)
	name := u.GetString(":name")
	if !ok {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		code, msg := models.DeleteUser(name)
		u.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
		}
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description 获取所有用户信息
// @router /all [get]
func (u *UserController) GetAllUsers() {
	token, err := u.ParseToken()
	if err != "" {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		u.ServeJSON()
		return
	}
	_, ok := token.Claims.(jwt.MapClaims)
	user, code, msg := models.GetAllUsers()
	if !ok {
		u.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		u.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
			"user": user,
		}
	}
	u.ServeJSON()
}

//验证token
func (u *UserController) ParseToken() (t *jwt.Token, err string) {
	authString := u.Ctx.Input.Header("Authorization")
	if authString == "" {
		return t, "token失效"
	}
	token, e := jwt.Parse(authString, func(token *jwt.Token) (interface{}, error) {
		return []byte("password"), nil
	})
	if e != nil {
		return token, "token失效"
	}
	if !token.Valid {
		return token, "token失效"
	}
	return token, ""
}

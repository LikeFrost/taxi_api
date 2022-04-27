package controllers

import (
	"encoding/json"
	"taxi_api/models"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type MoneyController struct {
	beego.Controller
}

// @Title AddMoney
// @Description 创建收支信息
// @router /add [post]
func (m *MoneyController) AddMoney() {
	token, err := m.ParseToken()
	if err != "" {
		m.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		var data map[string]interface{}
		json.Unmarshal(m.Ctx.Input.RequestBody, &data)
		types := data["types"].(string)
		in_time := data["in_time"].(string)
		in_type := data["in_type"].(string)
		in_money := data["in_money"].(string)
		people := data["people"].(string)
		more := data["more"].(string)
		if !ok {
			m.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.AddMoney(types, in_time, in_type, in_money, people, more)
			m.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	m.ServeJSON()
}

// @Title DeleteMoney
// @Description 删除收支信息
// @router /:money [delete]
func (m *MoneyController) DeleteMoney() {
	token, err := m.ParseToken()
	if err != "" {
		m.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		id, _ := m.GetInt(":money")
		if !ok {
			m.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.DeleteMoney(id)
			m.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	m.ServeJSON()
}

// @Title GetAllMoney
// @Description 获取全部收支信息
// @router /all [get]
func (m *MoneyController) GetAllMoney() {
	token, err := m.ParseToken()
	if err != "" {
		m.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			m.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg, money := models.GetAllMoney()
			m.Data["json"] = map[string]interface{}{
				"code":  code,
				"msg":   msg,
				"money": money,
			}
		}
	}
	m.ServeJSON()
}

//验证token
func (m *MoneyController) ParseToken() (t *jwt.Token, err string) {
	authString := m.Ctx.Input.Header("Authorization")
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

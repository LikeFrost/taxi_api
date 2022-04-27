package controllers

import (
	"encoding/json"
	"taxi_api/models"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type TrafficController struct {
	beego.Controller
}

// @Title AddTraffic
// @Description 创建违章信息
// @router /add [post]
func (t *TrafficController) AddTraffic() {
	token, err := t.ParseToken()
	if err != "" {
		t.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		var data map[string]interface{}
		json.Unmarshal(t.Ctx.Input.RequestBody, &data)
		beego.Error(data)
		car_id := data["car_id"].(string)
		user := data["user"].(string)
		content := data["content"].(string)
		time := data["time"].(string)
		location := data["location"].(string)
		money := data["money"].(string)
		score := data["score"].(string)
		state := data["state"].(string)
		if !ok {
			t.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.AddTraffic(car_id, user, content, time, location, money, score, state)
			t.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	t.ServeJSON()
}

// @Title GetTrafficByCarId
// @Description 通过车牌号获取违章
// @router /:car [get]
func (t *TrafficController) GetTrafficByCarId() {
	token, err := t.ParseToken()
	if err != "" {
		t.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		car_id := t.GetString(":car")
		if !ok {
			t.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg, traffic := models.GetTrafficByCarId(car_id)
			t.Data["json"] = map[string]interface{}{
				"code":     code,
				"msg":      msg,
				"traffics": traffic,
			}
		}
	}
	t.ServeJSON()
}

// @Title DeleteTraffic
// @Description 删除违章
// @router /:id [delete]
func (t *TrafficController) DeleteTraffic() {
	token, err := t.ParseToken()
	if err != "" {
		t.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		id, _ := t.GetInt(":id")
		if !ok {
			t.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.DeleteTraffic(id)
			t.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	t.ServeJSON()
}

// @Title GetAllTraffic
// @Description 获取所有违章
// @router /all [get]
func (t *TrafficController) GetAllTraffic() {
	token, err := t.ParseToken()
	if err != "" {
		t.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			t.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg, traffic := models.GetAllTraffic()
			t.Data["json"] = map[string]interface{}{
				"code":     code,
				"msg":      msg,
				"traffics": traffic,
			}
		}
	}
	t.ServeJSON()
}

//验证token
func (ty *TrafficController) ParseToken() (t *jwt.Token, err string) {
	authString := ty.Ctx.Input.Header("Authorization")
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

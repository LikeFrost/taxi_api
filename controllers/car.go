package controllers

import (
	"encoding/json"
	"taxi_api/models"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type CarController struct {
	beego.Controller
}

// @Title AddCar
// @Description 创建车辆信息
// @router /add [post]
func (c *CarController) AddCar() {
	token, err := c.ParseToken()
	if err != "" {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		var data map[string]interface{}
		json.Unmarshal(c.Ctx.Input.RequestBody, &data)
		beego.Error(data)
		car_id := data["car_id"].(string)
		color := data["color"].(string)
		car_type := data["car_type"].(string)
		id := data["id"].(string)
		count_money := data["count_money"].(string)
		state := data["state"].(string)
		ensure_company := data["ensure_company"].(string)
		ensure_year := data["ensure_year"].(string)
		safe_day := data["safe_day"].(string)
		ensure_day := data["ensure_year"].(string)
		if !ok {
			c.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.AddCar(car_id, color, car_type, id, count_money, state, ensure_company, ensure_year, safe_day, ensure_day)
			c.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	c.ServeJSON()
}

// @Title UpdateCar
// @Description 更新车辆信息
// @router / [post]
func (c *CarController) UpdateCar() {
	token, err := c.ParseToken()
	if err != "" {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		var data map[string]interface{}
		json.Unmarshal(c.Ctx.Input.RequestBody, &data)
		beego.Error(data)
		car_id := data["car_id"].(string)
		color := data["color"].(string)
		car_type := data["car_type"].(string)
		id := data["id"].(string)
		count_money := data["count_money"].(string)
		state := data["state"].(string)
		ensure_company := data["ensure_company"].(string)
		ensure_year := data["ensure_year"].(string)
		safe_day := data["safe_day"].(string)
		ensure_day := data["ensure_day"].(string)
		if !ok {
			c.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.UpdateCar(car_id, color, car_type, id, count_money, state, ensure_company, ensure_year, safe_day, ensure_day)
			c.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	c.ServeJSON()
}

// @Title GetCar
// @Description 获取车辆信息
// @router /:car [get]
func (c *CarController) GetCar() {
	token, err := c.ParseToken()
	if err != "" {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		car_id := c.GetString(":car")
		if !ok {
			c.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg, car := models.GetCar(car_id)
			c.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
				"car":  car,
			}
		}
	}
	c.ServeJSON()
}

// @Title DeleteCar
// @Description 删除车辆信息
// @router /:car [delete]
func (c *CarController) DeleteCar() {
	token, err := c.ParseToken()
	if err != "" {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		car_id := c.GetString(":car")
		if !ok {
			c.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.DeleteCar(car_id)
			c.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	c.ServeJSON()
}

// @Title GetAllCar
// @Description 获取全部车辆信息
// @router /all [get]
func (c *CarController) GetAllCar() {
	token, err := c.ParseToken()
	if err != "" {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg, cars := models.GetAllCar()
			c.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
				"cars": cars,
			}
		}
	}
	c.ServeJSON()
}

// @Title GetSearchCar
// @Description 搜索车辆信息
// @router /search/:search [get]
func (c *CarController) GetSearchCar() {
	token, err := c.ParseToken()
	if err != "" {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		c.ServeJSON()
		return
	}
	_, ok := token.Claims.(jwt.MapClaims)
	search := c.GetString(":search")
	car, code, msg := models.GetSearchCar(search)
	if !ok {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
			"cars": car,
		}
	}
	c.ServeJSON()
}

//验证token
func (c *CarController) ParseToken() (t *jwt.Token, err string) {
	authString := c.Ctx.Input.Header("Authorization")
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

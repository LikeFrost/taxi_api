package controllers

import (
	"encoding/json"
	"taxi_api/models"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type QuestionController struct {
	beego.Controller
}

// @Title AddQuestion
// @Description 创建投诉信息
// @router /add [post]
func (q *QuestionController) AddQuestion() {
	token, err := q.ParseToken()
	if err != "" {
		q.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		var data map[string]interface{}
		json.Unmarshal(q.Ctx.Input.RequestBody, &data)
		car_id := data["car_id"].(string)
		happen_time := data["happen_time"].(string)
		create_time := data["create_time"].(string)
		method := data["method"].(string)
		content := data["content"].(string)
		result := data["result"].(string)
		if !ok {
			q.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.AddQuestion(car_id, happen_time, create_time, method, content, result)
			q.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	q.ServeJSON()
}

// @Title GetQuestionByCarId
// @Description 通过车牌号获取投诉
// @router /:car [get]
func (q *QuestionController) GetQuestionByCarId() {
	token, err := q.ParseToken()
	if err != "" {
		q.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		car_id := q.GetString(":car")
		if !ok {
			q.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg, question := models.GetQuestionByCarId(car_id)
			q.Data["json"] = map[string]interface{}{
				"code":      code,
				"msg":       msg,
				"questions": question,
			}
		}
	}
	q.ServeJSON()
}

// @Title DeleteQuestion
// @Description 删除投诉
// @router /:id [delete]
func (q *QuestionController) DeleteQuestion() {
	token, err := q.ParseToken()
	if err != "" {
		q.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		id, _ := q.GetInt(":id")
		if !ok {
			q.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg := models.DeleteQuestion(id)
			q.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	q.ServeJSON()
}

// @Title GetAllQuestion
// @Description 获取所有投诉
// @router /all [get]
func (q *QuestionController) GetAllQuestion() {
	token, err := q.ParseToken()
	if err != "" {
		q.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			q.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			code, msg, question := models.GetAllQuestion()
			q.Data["json"] = map[string]interface{}{
				"code":      code,
				"msg":       msg,
				"questions": question,
			}
		}
	}
	q.ServeJSON()
}

// @Title GetSearchQuestion
// @Description 搜索投诉信息
// @router /search/:search [get]
func (q *QuestionController) GetSearchQuestion() {
	token, err := q.ParseToken()
	if err != "" {
		q.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
		q.ServeJSON()
		return
	}
	_, ok := token.Claims.(jwt.MapClaims)
	search := q.GetString(":search")
	question, code, msg := models.GetSearchQuestion(search)
	if !ok {
		q.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		q.Data["json"] = map[string]interface{}{
			"code":      code,
			"msg":       msg,
			"questions": question,
		}
	}
	q.ServeJSON()
}

//验证token
func (q *QuestionController) ParseToken() (t *jwt.Token, err string) {
	authString := q.Ctx.Input.Header("Authorization")
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

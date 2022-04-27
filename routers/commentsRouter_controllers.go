package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["taxi_api/controllers:CarController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:CarController"],
        beego.ControllerComments{
            Method: "UpdateCar",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:CarController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:CarController"],
        beego.ControllerComments{
            Method: "GetCar",
            Router: `/:car`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:CarController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:CarController"],
        beego.ControllerComments{
            Method: "DeleteCar",
            Router: `/:car`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:CarController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:CarController"],
        beego.ControllerComments{
            Method: "AddCar",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:CarController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:CarController"],
        beego.ControllerComments{
            Method: "GetAllCar",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:QuestionController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:QuestionController"],
        beego.ControllerComments{
            Method: "GetQuestionByCarId",
            Router: `/:car`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:QuestionController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:QuestionController"],
        beego.ControllerComments{
            Method: "DeleteQuestion",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:QuestionController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:QuestionController"],
        beego.ControllerComments{
            Method: "AddQuestion",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:QuestionController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:QuestionController"],
        beego.ControllerComments{
            Method: "GetAllQuestion",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:TrafficController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:TrafficController"],
        beego.ControllerComments{
            Method: "GetTrafficByCarId",
            Router: `/:car`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:TrafficController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:TrafficController"],
        beego.ControllerComments{
            Method: "DeleteTraffic",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:TrafficController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:TrafficController"],
        beego.ControllerComments{
            Method: "AddTraffic",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:TrafficController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:TrafficController"],
        beego.ControllerComments{
            Method: "GetAllTraffic",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:UserController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:UserController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:UserController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/:name`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:UserController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserByName",
            Router: `/:name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:UserController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAllUsers",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:UserController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:UserController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "LogUp",
            Router: `/logup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["taxi_api/controllers:UserController"] = append(beego.GlobalControllerRouter["taxi_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUserByName",
            Router: `/updateother`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

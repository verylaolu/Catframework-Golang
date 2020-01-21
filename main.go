package main

import (
	"github.com/gorilla/context"
	"net/http"
	"sso/Common"
	"sso/Route"
)


func main()  {

	db := Common.ConnectToDB()
	redis := Common.ConnectToRedis()

	// 使用db
	Common.SetDB(db)
	Common.SetRedis(redis)
	// go语法defer在整体函数运行完后执行
	defer db.Close()
	Route.Init()
	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
	//
	// user := UserInfo.User{Id:11,Username:"11"}
	// user.Age=12
	// user.Address="infosda"
	//
	//claim := jwt.MapClaims{
	//	"id":       user.Id,
	//	"username": user.Username,
	//	"age":		user.Age,
	//	"address":	user.Address,
	//	"nbf":      time.Now().Unix(),
	//	"iat":      time.Now().Unix(),
	//}
	//
	//tokenstr  := Common.EncodeJwtToken(claim)
	//fmt.Println(tokenstr)
	//
	//data :=Common.DecodeJwtToken(tokenstr)
	//info := UserInfo.User{}
	//info.Id =int(data["id"].(float64))
	//info.Username = data["username"].(string)
	//info.Age = int(data["age"].(float64))
	//info.Address = data["address"].(string)
	//fmt.Println(info.Id,info.Username,info.Age,info.Address)
	//

}

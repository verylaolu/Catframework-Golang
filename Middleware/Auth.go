package Middleware

import (
	"log"
	"net/http"
	"sso/Common"
)

const Authkey_pub  = "123321"

func Auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authkey := Common.PostValue(request,"key")
		if authkey != Authkey_pub {
			res :=Common.Result{500,"auth key is error",map[string]interface{}{"authkey":authkey,"bbb":"cccc"}}
			log.Println("error:", authkey)
			Common.Json_return(writer,res)

		}else{
			log.Println("success:", authkey)
			handlerFunc.ServeHTTP(writer, request)
		}
	}
}
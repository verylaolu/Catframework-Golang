package Controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"sso/App/Model"
	"sso/Common"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request){
	Common.Json_return(w,Common.Result{200,"api is error",map[string]interface{}{}})
	return
}
func Encodetoken(w http.ResponseWriter, r *http.Request){
	open_id := Common.PostValue(r,"open_id")
	userinfo,err := Model.GetUserByEmpno(Common.StrDecode(open_id))
	if(err != nil){
		Common.Json_return(w,Common.Result{500,"user is null", map[string]interface{}{}})
		return
	}
	claims := jwt.MapClaims{
			"openid": 		open_id,
			"empno":       userinfo.Empno,
			"title":		userinfo.Title,
			"email":	userinfo.Email,
			"time":      time.Now().Unix(),
		}
	jwt_token := Common.EncodeJwtToken(claims)
	Common.Json_return(w,Common.Result{200,"Encodetoken Success",map[string]interface{}{"token":jwt_token}})
	return
}

func Decodetoken(w http.ResponseWriter, r *http.Request)  {
	jwt_token := Common.PostValue(r,"token")
	jwt_info,err  := Common.DecodeJwtToken(jwt_token)
	if(err != nil){
		Common.Json_return(w,Common.Result{500,"jwt token is error token", map[string]interface{}{}})
		return
	}
	fmt.Println(err)
	userinfo := Model.User{}
	userinfo.Empno = jwt_info["empno"].(string)
	userinfo.Title = jwt_info["title"].(string)
	userinfo.Email = jwt_info["email"].(string)
	encodetime := int(jwt_info["time"].(float64))

	db_userinfo,err := Model.GetUserByEmpno(userinfo.Empno)
	if(db_userinfo.Empno == "" || err != nil){
		Common.Json_return(w,Common.Result{500,"user is null", map[string]interface{}{}})
		return
	}
	NowTime := time.Now().Unix()
	if int(NowTime) - int(encodetime) > 3600{
		Common.Json_return(w,Common.Result{500,"jwt token timeout", map[string]interface{}{}})
		return
	}else{
		Common.Json_return(w,Common.Result{200,"success", map[string]interface{}{"openid":Common.StrEncode(userinfo.Empno),"empno":userinfo.Empno,"username":userinfo.Username,"title":userinfo.Title,"email":userinfo.Email}})
		return
	}
}
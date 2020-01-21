package Common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)
const (
	SecretKey = "123321"
)

func EncodeJwtToken(claim jwt.MapClaims)string{
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	if tokenString, err := token.SignedString([]byte(SecretKey)); err == nil {
		return tokenString
	} else {
		return ""
	}
}
func secret()jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey),nil
	}
}
func DecodeJwtToken(tokenss string)(claim jwt.MapClaims,err error) {
	token, err := jwt.Parse(tokenss, secret())
	if err != nil {
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	return claim,err
}
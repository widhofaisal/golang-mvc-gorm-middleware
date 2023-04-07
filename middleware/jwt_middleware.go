package middleware

import (
	"learn/mvc/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, name string) (string, error) {
	
	// payload
	claims := jwt.MapClaims{}
	claims["userId"]=userId
	claims["name"]=name
	claims["exp"]=time.Now().Add(time.Hour*1).Unix()
	
	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// return + signature
	return token.SignedString([]byte(constant.SECRET_JWT))


}
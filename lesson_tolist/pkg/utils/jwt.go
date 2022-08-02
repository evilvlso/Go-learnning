package utils

import (
	"errors"
	"example.com/to_list/record"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const TokenExpireDuration = time.Hour*60
var SignMethod = jwt.SigningMethodHS256
var TokenSecret = []byte("vdail")

type User struct {
	Uid uint `json:"uid"`
}

type Claim struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(user User) string  {
	expireTime:=time.Now().Add(TokenExpireDuration)
	claim:=Claim{
		User:user,
		RegisteredClaims:jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		Issuer: "ToDoList",
		}}
	token:=jwt.NewWithClaims(SignMethod,claim)
	strToken,err:=token.SignedString(TokenSecret)
	if err != nil {
		record.Logger.Printf("GenerateToken Failed")
		return ""
	}
	return strToken
}

func CheckToken(token string) (*Claim,error) {
	tokenClaims,err:=jwt.ParseWithClaims(token,&Claim{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret,nil
	})
	if err != nil {
		return nil,err
	}
	if claim,ok:=tokenClaims.Claims.(*Claim);ok && tokenClaims.Valid{
		return claim,nil
	}
	return nil,errors.New("cant parse token")
}
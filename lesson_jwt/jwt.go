package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)
const TokenExpireDuration = time.Minute*3
var SignMethod = jwt.SigningMethodHS256
var TokenSecret = []byte("vdail")

type User struct {
	Uid uint `json:"uid"`
}

type Claims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}

//ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func main() {
	token:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IlVpZCI6MTd9LCJpc3MiOiJUb0RvTGlzdCIsImV4cCI6MTY1OTMxODUwNH0.6T48P-kjxbeXTM0Q68Af-vg2EeBcNe5Nqv0G5ljVAEg"
	claims,err:=ParseToken(token)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(claims)
}

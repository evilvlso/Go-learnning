package middleware

import (
	e "example.com/to_list/pkg/exceptions"
	"example.com/to_list/pkg/utils"
	"example.com/to_list/serialize"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		token:=c.GetHeader("X-Access")
		claims,err:=utils.CheckToken(token)
		if err!=nil{
			c.JSON(200,serialize.Response{
				Status: e.ErrorAuthCheckTokenFail,
				Msg: err.Error(),
			})
			c.Abort()
		}
		uid:=claims.User.Uid
		c.Set("uid",uid)
		c.Next()
	}
}
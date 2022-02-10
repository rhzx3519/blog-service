/**
@author ZhengHao Lou
Date    2022/02/07
*/
package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/go-programming-tour-book/blog-service/pkg/app"
    "github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var (
            token string
            ecode = errcode.Success
        )
        if s, exist := c.GetQuery("token"); exist {
            token = s
        } else {
            token = c.GetHeader("token")
        }
        if token == " " {
            ecode = errcode.InvalidParams
        } else {
            _, err := app.ParseToken(token)
            if err != nil {
                ecode = errcode.UnauthorizedTokenError
            }
        }

        if ecode != errcode.Success {
            response := app.NewResponse(c)
            response.ToErrorResponse(ecode)
            c.Abort()
            return
        }
        c.Next()
    }
}

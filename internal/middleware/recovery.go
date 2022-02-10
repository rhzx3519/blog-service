/**
@author ZhengHao Lou
Date    2022/02/09
*/
package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-programming-tour-book/blog-service/global"
    "github.com/go-programming-tour-book/blog-service/pkg/app"
    "github.com/go-programming-tour-book/blog-service/pkg/email"
    "github.com/go-programming-tour-book/blog-service/pkg/errcode"
    "time"
)

func Recovery() gin.HandlerFunc {
    defailtMailer := email.NewEmail(&email.SMTPInfo{
        Host:     global.EmailSetting.Host,
        Port:     global.EmailSetting.Port,
        IsSSL:    global.EmailSetting.IsSSL,
        UserName: global.EmailSetting.UserName,
        Password: global.EmailSetting.Password,
        From:     global.EmailSetting.From,
    })
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                s := "panic recover err: %v"
                global.Logger.WithCallersFrames().Errorf(s, err)

                err := defailtMailer.SendMail(global.EmailSetting.To,
                    fmt.Sprintf("Exception threw, occured at: %d", time.Now().Unix()),
                    fmt.Sprintf("err info: %v", err))
                if err != nil {
                    global.Logger.Fatalf("mail.SendMail err: %v", err)
                }
                app.NewResponse(c).ToErrorResponse(errcode.ServerError)
                c.Abort()
            }
        }()
        c.Next()
    }
}

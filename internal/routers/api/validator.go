/**
@author ZhengHao Lou
Date    2022/02/05
*/
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

func Validate(c *gin.Context, param interface{}) *errcode.Error {
	if valid, errs := app.BindAndValid(c, param); !valid {
		global.Logger.Errorf("app.BindAndValid err: %v\n", errs)
		return errcode.InvalidParams.WithDetails(errs.Errors()...)
	}
	return nil
}

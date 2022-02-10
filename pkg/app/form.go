/**
@author ZhengHao Lou
Date    2022/02/04
*/
package app

import (
    "github.com/gin-gonic/gin"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
)

type ValidError struct {
    Key     string
    Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
    return v.Message
}

func (v ValidErrors) Errors() []string {
    var errs []string
    for _, err := range v {
        errs = append(errs, err.Error())
    }
    return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
    var errs ValidErrors
    if err := c.ShouldBind(v); err != nil {
        v := c.Value("trans")
        trans, _ := v.(ut.Translator)
        if verrs, ok := err.(validator.ValidationErrors); ok {
            for key, value := range verrs.Translate(trans) {
                errs = append(errs, &ValidError{
                    Key:     key,
                    Message: value,
                })
            }
        }
        return false, errs
    }
    return true, nil
}

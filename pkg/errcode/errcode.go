/**
@author ZhengHao Lou
Date    2022/02/03
*/
package errcode

import (
    "fmt"
    "net/http"
)

type Error struct {
    code    int      `json:"code"`
    msg     string   `json:"msg"`
    details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
    if _, ok := codes[code]; ok {
        panic(fmt.Sprintf("error code %d is already exist", code))
    }
    codes[code] = msg
    return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
    return fmt.Sprintf("error code: %d, info: %s", e.code, e.msg)
}

func (e *Error) Code() int {
    return e.code
}

func (e *Error) Msg() string {
    return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
    return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
    return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
    newError := *e
    newError.details = []string{}
    for _, d := range details {
        newError.details = append(newError.details, d)
    }
    return &newError
}

func (e *Error) StatusCode() int {
    switch e.code {

    }
    return http.StatusInternalServerError
}








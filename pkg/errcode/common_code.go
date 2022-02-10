/**
@author ZhengHao Lou
Date    2022/02/03
*/
package errcode

var (
    Success                   = NewError(0, "success")
    ServerError               = NewError(10000000, "internal server err")
    InvalidParams             = NewError(10000001, "invalid params")
    NotFound                  = NewError(10000002, "not found")
    UnauthorizedAuthNotExist  = NewError(10000003, "authorize failed, cannot find AppKey and AppSecret")
    UnauthorizedTokenError    = NewError(10000004, "authorize failed, token err")
    UnauthorizedTokenTimeout  = NewError(10000005, "authorize failed, token timeout")
    UnauthorizedTokenGenerate = NewError(10000006, "authorize failed, token generate failed")
    TooManyRequests           = NewError(10000007, "too many requests")
)

var (
    ErrorGetTagListFail = NewError(20010001, "fetch tag list failed")
    ErrorCreateTagFail  = NewError(20010002, "create tag failed")
    ErrorUpdateTagFail  = NewError(20010003, "update tag failed")
    ErrorDeleteTagFail  = NewError(20010004, "delete tag failed")
    ErrorCountTagFail   = NewError(20010005, "count tag failed")
)

/**
@author ZhengHao Lou
Date    2022/02/07
*/
package dao

import "github.com/go-programming-tour-book/blog-service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}

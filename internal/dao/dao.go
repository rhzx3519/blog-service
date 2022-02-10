/**
@author ZhengHao Lou
Date    2022/02/04
*/
package dao

import "github.com/jinzhu/gorm"

type Dao struct {
    engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
    return &Dao{engine}
}

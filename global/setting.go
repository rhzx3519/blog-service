/**
@author ZhengHao Lou
Date    2022/02/03
*/
package global

import (
    "github.com/go-programming-tour-book/blog-service/pkg/logger"
    "github.com/go-programming-tour-book/blog-service/pkg/setting"
    "github.com/jinzhu/gorm"
)

var (
    DBEngine *gorm.DB

    Logger *logger.Logger

    ServerSetting   *setting.ServerSettings
    AppSetting      *setting.AppSettings
    DatabaseSetting *setting.DatabaseSettings
    JwtSetting      *setting.JwtSettings
    EmailSetting    *setting.EmailSettings
)

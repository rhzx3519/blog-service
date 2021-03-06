/**
@author ZhengHao Lou
Date    2022/02/03
*/
package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type ArticleSwagger struct {
    List  *[]Article
    Pager *app.Pager
}

type Article struct {
    *Model
    Title          string `json:"title"`
    Desc           string `json:"desc"`
    Content        string `json:"content"`
    ConverImageUrl string `json:"conver_image_url"`
    State          uint8  `json:"state"`
}

func (a Article) TableName() string {
    return "blog_article"
}

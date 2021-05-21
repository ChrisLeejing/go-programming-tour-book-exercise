package dao

import "go-programming-tour-book-exercise/chapter-2-blog-service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.GetAuth(d.engine)
}

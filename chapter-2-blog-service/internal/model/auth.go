package model

import (
	"github.com/jinzhu/gorm"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
)

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) GetAuth(db *gorm.DB) (Auth, error) {
	var auth Auth
	err := db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		global.Logger.Errorf("authentication fail, err: %v", err)
		return auth, err
	}

	return auth, nil
}

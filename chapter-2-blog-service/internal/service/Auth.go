package service

import (
	"errors"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/model"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service/validate"
)

func (svc *Service) GetAuth(param *validate.GetAuthRequest) (model.Auth, error) {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return auth, err
	}

	if auth.ID > 0 {
		return auth, nil
	}

	return model.Auth{}, errors.New("Auth info not found. ")
}

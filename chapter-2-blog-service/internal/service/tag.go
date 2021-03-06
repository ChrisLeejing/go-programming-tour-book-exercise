package service

import (
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/model"
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/service/validate"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/errcode"
)

func (svc *Service) CountTag(param *validate.CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *validate.ListTagRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *validate.CreateTagRequest) error {
	// tag is existed.
	tag, _ := svc.dao.GetTagByName(param.Name)
	if tag.ID > 0 && tag.Name != "" {
		return errcode.ErrorTagExisted
	}

	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *validate.UpdateTagRequest) error {
	// add GetTagById() logic
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *validate.DeleteTagRequest) error {
	// add GetTagById() logic
	return svc.dao.DeleteTag(param.ID)
}

func (svc *Service) GetTagById(param *validate.GetTagByIdRequest) (model.Tag, error) {
	return svc.dao.GetTagById(param.ID)
}

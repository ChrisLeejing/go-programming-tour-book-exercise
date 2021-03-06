package dao

import (
	"go-programming-tour-book-exercise/chapter-2-blog-service/internal/model"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
)

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	// query: no need for model.Tag.Model
	tag := model.Tag{Name: name, State: state}
	// pagination.go: page -> pageOffset
	pageOffSet := app.GetPageOffSet(page, pageSize)
	return tag.List(d.engine, pageOffSet, pageSize)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createBy,
		},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	// fix the bug: the gorm does not update the zero-field.
	values := map[string]interface{}{
		"State":      state,
		"ModifiedBy": modifiedBy,
	}
	if values["Name"] != "" {
		values["Name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	return tag.Delete(d.engine)
}

func (d *Dao) GetTagById(id uint32) (model.Tag, error) {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	return tag.GetTagById(d.engine)
}

func (d *Dao) GetTagByName(name string) (model.Tag, error) {
	tag := model.Tag{
		Name: name,
	}

	return tag.GetTagByName(d.engine)
}

func (d *Dao) GetTagByIds(ids []uint32) ([]model.Tag, error) {
	tag := model.Tag{}

	return tag.GetTagByIds(d.engine, ids)
}

package model

import (
	"github.com/jinzhu/gorm"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	list  []*Tag
	pager *app.Pager
}

func (Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// reference: https://gorm.io/docs/query.html
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db.Where("state = ?", t.State)
	// SELECT * FROM `tag`  WHERE (is_del = 0) LIMIT 5 OFFSET 5
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// reference: https://gorm.io/docs/create.html
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// reference: https://gorm.io/docs/update.html
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0).Update(values).Error
}

// reference: https://gorm.io/docs/delete.html
func (t Tag) Delete(db *gorm.DB) error {
	// soft delete.
	// t.IsDel = 1
	// return t.Update(db)
	return db.Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error // really delete.
}

func (t Tag) GetTagById(db *gorm.DB) (Tag, error) {
	err := db.Where("id = ? AND is_del = ?", t.ID, 0).First(&t).Error

	return t, err
}

func (t Tag) GetTagByName(db *gorm.DB) (Tag, error) {
	err := db.Where("name = ? AND is_del = ?", t.Name, 0).First(&t).Error

	return t, err
}

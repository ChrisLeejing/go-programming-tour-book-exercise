package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `gorm:"type:longtext" json:"context"`
	State         uint8  `json:"state"`
}

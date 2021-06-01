package validate

type Article struct {
	ID          uint32   `json:"id"`
	Title       string   `json:"title"`
	Desc        string   `json:"desc"`
	CoverImgUrl string   `json:"cover_img_url"`
	Content     string   `json:"content"`
	CreatedBy   string   `json:"created_by"`
	ModifiedBy  string   `json:"modified_by"`
	State       uint8    `json:"state"`
	TagIds      []uint32 `json:"tag_ids"`
}

type CreateArticleRequest struct {
	Title       string   `form:"title" json:"title" binding:"required,max=100"`
	Desc        string   `json:"desc"`
	CoverImgUrl string   `json:"cover_img_url"`
	Content     string   `json:"content"`
	CreatedBy   string   `form:"created_by" json:"created_by" binding:"required,min=3,max=100"`
	State       uint8    `form:"state,default=1" json:"state" binding:"oneof=0 1"`
	TagIds      []uint32 `json:"tag_ids"`
}

type DeleteArticleRequest struct {
	// todo 20210517
}

type UpdateArticleRequest struct {
}

type ListArticleRequest struct {
}

type CountArticleRequest struct {
}

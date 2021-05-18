package validate

// https://github.com/gin-gonic/gin#model-binding-and-validation
// `form:"user" json:"user" xml:"user"  binding:"required"`
// alike validate layer: CountTagRequest, CreateTagRequest, ListTagRequest, UpdateTagRequest, DeleteTagRequest.
type CountTagRequest struct {
	Name  string `form:"name" json:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" json:"state" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" json:"name" binding:"max=100"`
	State     uint8  `form:"state,default=1" json:"state" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=3,max=100"`
}

type ListTagRequest struct {
	Name  string `form:"name" json:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" json:"state" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" json:"id" binding:"required,gte=1"`
	Name       string `form:"name" json:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" json:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" json:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" json:"id" binding:"required,gte=1"`
}

package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCountTagFail   = NewError(20010002, "获取标签总数失败")
	ErrorCreateTagFail  = NewError(20010003, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010004, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010005, "逻辑(物理)删除标签失败")
	ErrorGetTagFail     = NewError(20010006, "获取标签失败")
	ErrorTagExisted     = NewError(20010007, "该标签已存在")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")

	ErrorGetAuthFail = NewError(20040001, "获取验证失败")
)

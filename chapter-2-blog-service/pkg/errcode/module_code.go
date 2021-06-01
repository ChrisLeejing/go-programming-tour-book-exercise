package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCountTagFail   = NewError(20010002, "获取标签总数失败")
	ErrorCreateTagFail  = NewError(20010003, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010004, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010005, "逻辑(物理)删除标签失败")
	ErrorGetTagFail     = NewError(20010006, "获取标签失败")
	ErrorTagExisted     = NewError(20010007, "该标签已存在")
	ErrorGetTagsFail    = NewError(20010008, "获取多个标签失败")

	ErrorCreateArticleFail     = NewError(20020001, "创建文章失败")
	ErrorCreateArticleOnlyFail = NewError(20020002, "创建文章SQL执行失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")

	ErrorGetAuthFail = NewError(20040001, "获取验证失败")

	ErrorCreateArticleTagsFail = NewError(20060001, "创建文章&标签失败")
)

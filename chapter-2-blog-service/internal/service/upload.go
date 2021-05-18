package service

import (
	"errors"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	filename := upload.GetFileName(fileHeader.Filename)

	if !upload.CheckContainExt(fileType, filename) {
		return nil, errors.New("file extension not supported")
	}
	if upload.CheckExceedMaxSize(fileType, file) {
		return nil, errors.New("file exceeds max size limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		// os.ModePerm: Unix permission bits, 0o777
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("file directory failed to create")
		}
	}

	// comment blow code in Windows, open it in linux or Mac.
	// if !upload.CheckPermission(uploadSavePath) {
	// 	return nil, errors.New("insufficient file permissions")
	// }

	dst := uploadSavePath + "/" + filename
	accessUrl := global.AppSetting.UploadSavePath + "/" + filename
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	fileInfo := &FileInfo{Name: filename, AccessUrl: accessUrl}

	return fileInfo, nil
}

package upload

import (
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func CheckSavePath(dst string) bool {
	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	_, err := os.Stat(dst)

	// IsNotExist returns a boolean indicating whether the error is known to
	// report that a file or directory does not exist. It is satisfied by
	// ErrNotExist as well as some syscall errors.
	return os.IsNotExist(err)
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage, TypeExcel, TypeTxt:
		for _, v := range global.AppSetting.UploadImageAllowExts {
			if ext == strings.ToUpper(v) {
				return true
			}
		}
	}

	return false
}

func CheckExceedMaxSize(t FileType, file multipart.File) bool {
	bytes, _ := ioutil.ReadAll(file)
	size := len(bytes)
	switch t {
	case TypeImage:
		// UploadImageMaxSize: 5  # MB
		if size > global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

// func CheckPermission(dst string) bool {
// 	// Stat returns a FileInfo describing the named file.
// 	_, err := os.Stat(dst)
// 	return os.IsPermission(err)
// }

func CreateSavePath(dst string, perm os.FileMode) error {
	// A FileMode represents a file's mode and permission bits.
	if err := os.MkdirAll(dst, perm); err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	return nil
}

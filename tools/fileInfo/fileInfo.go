package fileInfo

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

func CalcMd5(fc *multipart.FileHeader) (md5v string, err error) {
	src, err := fc.Open()
	if err != nil {
		return
	}
	defer src.Close()

	md5h := md5.New()
	io.Copy(md5h, src)
	return fmt.Sprintf("%x", md5h.Sum([]byte(""))), nil
}

func CalcSize(file *multipart.FileHeader) int64 {
	return file.Size

}

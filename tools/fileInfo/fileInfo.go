package fileInfo

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
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

func CalcMd5FromPath(p string) string {

	fIn, err := os.Open(p)
	defer fIn.Close()
	if err != nil {
		zap.L().Error("文件打开失败", zap.Error(err))
		return ""
	}
	fReader := bufio.NewReader(fIn)
	md5h := md5.New()
	io.Copy(md5h, fReader)
	return fmt.Sprintf("%x", md5h.Sum([]byte("")))
}

func CalcSize(file *multipart.FileHeader) int64 {
	return file.Size

}

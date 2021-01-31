package logic

import (
	"go.uber.org/zap"
	"io/ioutil"
	"net/url"
	"os"
)

func Download(path string) (bs []byte, err error) {
	// 获取要返回的文件数据流
	// 看你文件存在哪里了，本地就直接os.Open就可以了，总之是要获取一个[]byte
	//去掉请求中的download
	//sff := strings.Split(path, "download")[1]
	////去掉字符串第一个字符
	//sf := sff[1:]

	//url中有中文，自动化转码
	sf, _ := url.PathUnescape(path[10:])
	fileContent, err := os.Open(sf)
	if err != nil {
		zap.L().Info(path)
		zap.L().Error("打开文件失败！", zap.Error(err))
		return
	}
	defer fileContent.Close()

	bs, err = ioutil.ReadAll(fileContent)
	if err != nil {
		zap.L().Error("文件读取失败！", zap.Error(err))
		return nil, err
	}

	return

}

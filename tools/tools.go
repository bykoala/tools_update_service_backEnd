package tools

import (
	"go.uber.org/zap"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
)

//获取当前程序所在服务器的ip
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

//目录不存在则创建
func Mkdir(p string) {

	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.Mkdir(p, os.ModePerm)
	}
}

func GetRootDir() string {
	_, fileStr, _, _ := runtime.Caller(1)
	return fileStr
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func CopyFile(srcFileName string, dstFileName string) {
	//打开源文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		zap.L().Error("源文件读取失败,原因是:", zap.Error(err))
		//log.Fatalf("源文件读取失败,原因是:%v\n", err)
	}
	defer func() {
		err = srcFile.Close()
		if err != nil {
			zap.L().Error("源文件关闭失败,原因是:", zap.Error(err))
			//log.Fatalf("源文件关闭失败,原因是:%v\n", err)
		}
	}()

	//创建目标文件,稍后会向这个目标文件写入拷贝内容
	distFile, err := os.Create(dstFileName)
	if err != nil {
		zap.L().Error("目标文件创建失败,原因是:", zap.Error(err))
		//log.Fatalf("目标文件创建失败,原因是:%v\n", err)
	}
	defer func() {
		err = distFile.Close()
		if err != nil {
			zap.L().Error("目标文件关闭失败,原因是:", zap.Error(err))
			//log.Fatalf("目标文件关闭失败,原因是:%v\n", err)
		}
	}()
	//定义指定长度的字节切片,每次最多读取指定长度
	var tmp = make([]byte, 1024*4)
	//循环读取并写入
	for {
		n, err := srcFile.Read(tmp)
		n, _ = distFile.Write(tmp[:n])
		if err != nil {
			if err == io.EOF { //读到了文件末尾,并且写入完毕,任务完成返回(关闭文件的操作由defer来完成)
				return
			} else {
				zap.L().Error("拷贝过程中发生错误,错误原因为:", zap.Error(err))
				//log.Fatalf("拷贝过程中发生错误,错误原因为:%v\n", err)
			}
		}
	}
}

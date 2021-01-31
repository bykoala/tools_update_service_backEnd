package tools

import (
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

package main

import (
	"os"
	"study-error-package-go-original/errorutil"
)

func main() {
	// 测试error封装方法
	_, err := os.Open("test.txt")
	errorutil.ErrorPanic(err, "打开文件失败")
	errorutil.ErrorPrint(err, "打开文件失败")

}

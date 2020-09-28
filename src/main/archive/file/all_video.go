package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
 * 查看文件夹下所有视频
 */
func main() {

	fileName := "F:\\ElasticSearch顶尖高手系列课程\\1_快速入门篇视频"

	if !strings.HasSuffix(fileName, "\\") {
		fileName += "\\"
	}

	stat, err := os.Stat(fileName)

	if err != nil {
		fmt.Println("[文件或文件夹] 不存在!")
		return
	}

	if stat.IsDir() {
		fmt.Println("当前文件夹：", fileName)
	} else {
		fmt.Println("当前文件：", fileName)
		return
	}

	eachDir(fileName)
}

/**
 * 遍历文件夹
 */
func eachDir(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			eachDir(path + fileName + "\\")
		} else {
			if strings.HasSuffix(fileName, ".mp4") || strings.HasSuffix(fileName, ".avi") {
				fmt.Println(path + fileName)
			}
		}
	}
}

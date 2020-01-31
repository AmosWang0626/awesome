package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

/*
go get golang.org/x/text
在Windows命令行下执行`chcp`, 输出936, 也即GBK
*/
func GbkToUtf8(src string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder())
	text, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("ERROR GbkToUtf8 :: ", err)
		return ""
	}
	return string(text)
}

func Utf8ToGbk(src string) string {
	reader := transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder())
	text, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("ERROR Utf8ToGbk :: ", err)
		return ""
	}
	return string(text)
}

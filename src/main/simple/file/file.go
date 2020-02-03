package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

func main() {

	//readFile1()
	//readFile2()
	//readAndWriteFile1()
	//copyContent()
	//pathOrFileExist()
	copyFile()

	fmt.Println("read file end.")

}

func copyFile() {
	fileName := "E:\\project\\golang\\awesome\\src\\main\\res\\text.txt"
	fileName2 := "E:\\project\\golang\\awesome\\src\\main\\res\\text2.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	reader := bufio.NewReader(file)

	file2, _ := os.OpenFile(fileName2, syscall.O_CREAT|syscall.O_RDWR, 0777)
	defer file2.Close()
	writer := bufio.NewWriter(file2)

	_, err := io.Copy(writer, reader)
	log.Fatalln(err)
}

func pathOrFileExist() {
	fileName := "E:\\project\\golang\\awesome\\src\\main\\res1"

	_, err := os.Stat(fileName)

	if err == nil {
		fmt.Println("[文件或文件夹] 存在")
		return
	} else if os.IsNotExist(err) {
		fmt.Println("[文件或文件夹] 不存在!")
	} else {
		fmt.Println("[文件或文件夹] 状态异常!")
		log.Fatalln(err)
	}

}

func copyContent() {
	fileName := "E:\\project\\golang\\awesome\\src\\main\\res\\text.txt"
	fileName2 := "E:\\project\\golang\\awesome\\src\\main\\res\\text2.txt"

	fileByte, _ := ioutil.ReadFile(fileName)

	fmt.Println(string(fileByte))

	_ = ioutil.WriteFile(fileName2, fileByte, 0666)
}

func readAndWriteFile1() {
	// OpenFile(name string, flag int, perm FileMode)
	// name:文件名字;
	// flag:文件打开模式 [只读: O_RDONLY] [只写: O_WRONLY] [读写: O_RDWR] [不存在就创建: O_CREAT] [追加: O_APPEND] [打开并清空: O_TRUNC (慎用)]
	// perm:文件权限,linux下有效[r 4; w 2; x 1]
	filePath := "E:\\project\\golang\\awesome\\src\\main\\res\\text.txt"
	file, err := os.OpenFile(filePath, syscall.O_WRONLY|syscall.O_TRUNC, 7)
	if err != nil {
		log.Fatalln(err)
		return
	}

	printFile(file)

	fmt.Println("...........................")

	writer := bufio.NewWriter(file)
	writer.WriteString("APPEND FINISH!\n")
	err = writer.Flush()
	if err != nil {
		log.Fatalln(err)
		return
	}

	readFile2()
}

func readFile2() {
	content, err := ioutil.ReadFile("E:\\project\\golang\\awesome\\src\\main\\res\\text.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(content))
}

func readFile1() {
	// file文件指针, 别名文件对象、文件句柄
	file, err := os.Open("E:\\project\\golang\\awesome\\src\\main\\res\\text.txt")
	defer FileClose(file)
	if err != nil {
		log.Fatalln(err)
		return
	}

	printFile(file)

}

func printFile(file *os.File) {
	fmt.Printf("FILE=%p, FILE_NAME=%v\n", file, file.Name())
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(str)
	}
}

func FileClose(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatalln("ERROR :: ", err)
	}
}

// os.OpenFile(name string, flag int, perm FileMode)
// >>> OpenFile 第二个参数 flag
//const (
//	// Invented values to support what package os expects.
//	O_RDONLY   = 0x00000 只读
//	O_WRONLY   = 0x00001 只写
//	O_RDWR     = 0x00002 读写
//	O_CREAT    = 0x00040 如果不存在就创建新的
//	O_EXCL     = 0x00080 和O_CREAT配合使用,文件必须不存在
//	O_NOCTTY   = 0x00100
//	O_TRUNC    = 0x00200 如果可能,打开时清空文件
//	O_NONBLOCK = 0x00800
//	O_APPEND   = 0x00400 写操作时将数据追加到文件尾部
//	O_SYNC     = 0x01000 打开文件用于同步I/O
//	O_ASYNC    = 0x02000
//	O_CLOEXEC  = 0x80000
//)
// >>> OpenFile 第三个参数 perm
// r 4; w 2; x 1 >>>>>>>>>>> Read Write eXecute

package fileMountain

import (
	"fmt"
	"io"
	"os"
)

// 使用文件中索引读取文件
// Q1 如果根据struct 预测存储的字节数？

const filePath = "./2.txt"

func getSizeOfStudent(filePath string) {
	// open file
	file, err := os.Open(filePath)
	errorHandler(err)
	defer file.Close()
	content, err := io.ReadAll(file)
	errorHandler(err)
	var start = 0
	for ; start < len(content); start++ {
		if content[start] == '\n' {
			break
		}
	}
	fmt.Println(start)
}

func loadIndex(filePath string) {

}

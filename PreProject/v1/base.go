/*
 * @Author: Coffeerr
 * @Date: 2022-08-16 11:05:16
 * @LastEditTime: 2022-08-17 10:51:55
 * @Description: 初始配置
 */
package v1

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/goombaio/namegenerator"
)

// var headerSize = uint32(10)
var offset = 0

func main() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name1 := nameGenerator.Generate()

	fmt.Println(name1)
	name2 := nameGenerator.Generate()
	fmt.Println(name2)
}
func errorHandler(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "something wrong:%v", err)
		os.Exit(1)
	}
}

/**
 * @description: 数据库文件初始化
 * @param {string} output
 * @param {int} size
 * @return {*}
 */
func initDataBase(output string, size int) {
	f, err := os.Create(output)
	errorHandler(err)
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	var offset = 0
	for i := 0; i < size; i++ {
		stu := &Student{}
		stu.index = uint16(i)
		stu.interest = []byte("bastetable")
		stu.name = []byte(nameGenerator.Generate())
		stu.interestSize = uint32(len(stu.interest))
		stu.nameSize = uint32(len(stu.name))
		// fmt.Fprintf(f, "%v", stu)
		buf, err := stu.Encode()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		f.WriteAt(buf, int64(offset))
		offset += int(stu.getSize())
	}
}

func read(input string) ([]*Student, error) {
	f, err := os.Open(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil, err
	}
	res := make([]*Student, 100)
	var count = 0
	for {
		buf := make([]byte, headerSize)
		_, err2 := f.ReadAt(buf, int64(offset))
		if err2 == io.EOF {
			break
		}
		stu, err := Decode(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return nil, err
		}
		offset += int(headerSize)
		name := make([]byte, stu.nameSize)
		_, err2 = f.ReadAt(name, int64(offset))
		if err2 == io.EOF {
			break
		}
		stu.name = name
		offset += int(stu.nameSize)
		interest := make([]byte, stu.interestSize)
		_, err2 = f.ReadAt(interest, int64(offset))
		if err2 == io.EOF {
			break
		}
		stu.interest = interest
		res[count] = stu
		count++
		offset += int(stu.interestSize)
	}
	return res, nil
}

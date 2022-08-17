/*
 * @Author: Coffeerr
 * @Date: 2022-08-16 11:05:50
 * @LastEditTime: 2022-08-17 10:55:26
 * @Description:
 */
package v1

import (
	"fmt"
	"testing"
)

var filePath = "./3.txt"

func Test_1(t *testing.T) {
	initDataBase(filePath, 10)
}

func Test_read(t *testing.T) {
	s, err := read(filePath)
	if err != nil {

	}
	fmt.Println(len(s))
	for i := 0; i < 100; i++ {
		stu := s[i]
		if stu == nil {
			break
		}
		fmt.Println(string(stu.name), "---", string(stu.interest))

	}
}

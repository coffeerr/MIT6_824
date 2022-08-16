package fileMountain

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	index int
	name  string
	age   int
}

func (stu Student) String() string {
	return fmt.Sprintf("Student id : %v \t age : %v \t name : %v\n", stu.index, stu.age, stu.name)
}

func initFile(num int) string {
	create, err := os.Create("./2.txt")
	filePath := create.Name()
	nameList := []string{"Jack", "Tom", "John"}
	errorHandler(err)

	for i := 0; i < num; i++ {
		index := rand.Int() % 3
		if _, err := fmt.Fprintf(create, "%d#%d#%s\n", i, i*i%100, nameList[index]); err != nil {
			errorHandler(err)
		}
	}
	create.Close()
	return filePath
}
func errorHandler(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "something wrong:%v", err)
		os.Exit(1)
	}
}
func open(dirpath string) {
	file, err := os.Open(dirpath)
	defer file.Close()
	errorHandler(err)
	all, err := io.ReadAll(file)
	errorHandler(err)
	s := string(all)
	split := strings.Split(s, "\n")
	for _, str := range split {
		if len(str) == 0 {
			break
		}
		conv := studentConv(str)
		fmt.Printf("%v", conv)
	}
}
func openWithBuffer(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	errorHandler(err)
	var buffer = make([]byte, 10)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				break
			}
		}
		errorHandler(err)
		fmt.Println(n)
	}

}

func studentConv(str string) Student {
	stu := strings.Split(str, "#")
	id, err := strconv.Atoi(stu[0])
	errorHandler(err)
	age, err := strconv.Atoi(stu[1])
	errorHandler(err)
	return Student{
		index: id,
		age:   age,
		name:  stu[2],
	}
}

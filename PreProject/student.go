package fileMountain

import "fmt"

type Student struct {
	index int
	name  string
	age   int
}

func (stu Student) String() string {
	return fmt.Sprintf("Student id : %v \t age : %v \t name : %v\n", stu.index, stu.age, stu.name)
}

package fileMountain

import (
	"testing"
)

// 这种方法的缺陷在于，没有办法利用到索引的结构
// 换句话说，没有办发根据索引快速定位数据的记录
// 现在需要对序列化方式进行改装，需要对数据库记录进行改写

func Test_Open(t *testing.T) {
	fileName := initFile(10)
	open(fileName)
	//fmt.Println(fileName)
}

func Test_Size_Count(t *testing.T) {
	getSizeOfStudent("./2.txt")
}

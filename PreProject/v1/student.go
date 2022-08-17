package v1

import (
	"encoding/binary"
)

var headerSize = uint32(10)

type Student struct {
	index        uint16
	nameSize     uint32
	interestSize uint32
	name         []byte
	interest     []byte
}

func (stu *Student) getSize() int64 {
	return (int64(headerSize) + int64(stu.nameSize) + int64(stu.interestSize))
}

func (stu *Student) Encode() ([]byte, error) {
	buf := make([]byte, stu.getSize())
	binary.BigEndian.PutUint16(buf[0:2], stu.index)
	binary.BigEndian.PutUint32(buf[2:6], stu.nameSize)
	binary.BigEndian.PutUint32(buf[6:10], stu.interestSize)
	copy(buf[headerSize:headerSize+stu.nameSize], stu.name)
	copy(buf[headerSize+stu.nameSize:headerSize+stu.nameSize+stu.interestSize], stu.interest)
	return buf, nil
}

func Decode(buf []byte) (*Student, error) {
	return &Student{
		index:        binary.BigEndian.Uint16(buf[0:2]),
		nameSize:     binary.BigEndian.Uint32(buf[2:6]),
		interestSize: binary.BigEndian.Uint32(buf[6:10]),
	}, nil
}

// 消息解析 Create By Ace 2013-2-1
package ace

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type AceBuffer struct {
	message []byte
	length  int
}

func NewBuffer(in []byte) *AceBuffer { return &AceBuffer{in, len(in)} }

/*读取int*/
func (ab *AceBuffer) ReadInt() int {
	in := ab.message[0:4]
	result := int(in[3]) | (int(in[2]) << 8) | (int(in[1]) << 16) | (int(in[0]) << 24)
	ab.message = ab.message[4:ab.length]
	ab.length = len(ab.message)
	return int(result)
}

/*读取float*/
func (ab *AceBuffer) ReadFloat() float32 {
	in := ab.message[0:4]
	var result float32
	buf := bytes.NewBuffer(in)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		fmt.Println("float解析失败", err)
	}
	ab.message = ab.message[4:ab.length]
	ab.length = len(ab.message)
	return result
}

/*读取double*/
func (ab *AceBuffer) ReadDouble() float64 {
	in := ab.message[0:8]
	var result float64
	buf := bytes.NewBuffer(in)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		fmt.Println("double解析失败", err)
	}
	ab.message = ab.message[8:ab.length]
	ab.length = len(ab.message)
	return result
}

func (ab *AceBuffer) ReadByte() byte {
	result := ab.message[0]
	ab.message = ab.message[1:ab.length]
	ab.length = len(ab.message)
	return result
}

/*读取string*/
func (ab *AceBuffer) ReadString(length int) string {
	in := ab.message[0:length]
	ab.message = ab.message[length:ab.length]
	ab.length = len(ab.message)
	return string(in)
}

/*写入int*/
func (ab *AceBuffer) WriteInt(value int) {
	b := make([]byte, 4)
	b[0] = byte(value >> 24)
	b[1] = byte(value >> 16)
	b[2] = byte(value >> 8)
	b[3] = byte(value)
	ab.message = append(ab.message, b...)
	ab.length = len(ab.message)
}

/*写入byte*/
func (ab *AceBuffer) WriteByte(value byte) {
	ab.message = append(ab.message, value)
	ab.length = len(ab.message)
}

/*写入float*/
func (ab *AceBuffer) WriteFloat(value float32) {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	err := binary.Write(buf, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("float写入失败", err)
	}
	ab.message = append(ab.message, buf.Bytes()...)
	ab.length = len(ab.message)
}

/*写入double*/
func (ab *AceBuffer) WriteDouble(value float64) {
	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	err := binary.Write(buf, binary.BigEndian, &value)
	if err != nil {
		fmt.Println("double写入失败", err)
	}
	ab.message = append(ab.message, buf.Bytes()...)
	ab.length = len(ab.message)
}

/*写入String*/
func (ab *AceBuffer) WriteString(value string) {
	ab.message = append(ab.message, []byte(value)...)
	ab.length = len(ab.message)
}

func (ab *AceBuffer) Bytes() []byte {
	return ab.message
}

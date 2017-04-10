package hessiancoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func packInt8(v int8) (r []byte, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, v); err != nil {
		return
	}
	r = buf.Bytes()
	return
}

// packInt16 [10].pack('n').bytes => [0, 10]
func packInt16(v int16) (r []byte, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, v); err != nil {
		return
	}
	r = buf.Bytes()
	return
}

// packUInt16 [10].pack('n').bytes => [0, 10]
func packUInt16(v uint16) (r []byte, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, v); err != nil {
		return
	}
	r = buf.Bytes()
	return
}

// packInt32 [10].pack('N').bytes => [0, 0, 0, 10]
func packInt32(v int32) (r []byte, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, v); err != nil {
		return
	}
	r = buf.Bytes()
	return
}

// packInt64 [10].pack('q>').bytes => [0, 0, 0, 0, 0, 0, 0, 10]
func packInt64(v int64) (r []byte, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, v); err != nil {
		return
	}
	r = buf.Bytes()
	return
}

// packFloat64 [10].pack('G').bytes => [64, 36, 0, 0, 0, 0, 0, 0]
func packFloat64(v float64) (r []byte, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, v); err != nil {
		return
	}
	r = buf.Bytes()
	return
}

// unpackInt16 (0,2).unpack('n')
func unpackInt16(b []byte) (pi int16, err error) {
	if err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi); err != nil {
		return
	}
	return
}

// unpackInt32 (0,4).unpack('N')
func unpackInt32(b []byte) (pi int32, err error) {
	if err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi); err != nil {
		return
	}
	return
}

// unpackInt64 long (0,8).unpack('q>')
func unpackInt64(b []byte) (pi int64, err error) {
	if err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi); err != nil {
		return
	}
	return
}

// unpackFloat64 Double (0,8).unpack('G)
func unpackFloat64(b []byte) (pi float64, err error) {
	if err = binary.Read(bytes.NewReader(b), binary.BigEndian, &pi); err != nil {
		return
	}
	return
}

// 将字节数组格式化成 hex
func SprintHex(b []byte) (rs string) {
	rs = fmt.Sprintf("[]byte{")
	for _, v := range b {
		rs += fmt.Sprintf("0x%02x,", v)
	}
	rs = strings.TrimSpace(rs)
	rs += fmt.Sprintf("}\n")
	return
}

// HostCheck make host conforms the HTTP request
func HostCheck(host string) string {
	index := strings.Index(host, "http://")
	if index != -1 {
		return host
	}
	return "http://" + host
}

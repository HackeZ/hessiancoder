package hessiancoder

import (
	"log"
	"runtime"
)

type Encoder struct {
}

const (
	ChunkSize = 0x8000 // 64k each chunk
)

func init() {
	// for debug
	_, filename, _, _ := runtime.Caller(1)
	log.SetPrefix(filename + "\n")
}

// encodeBinary
// binary ::= [x62(b)]  b1 b0 <binary-data> binary  # non-final chunk
//        ::= [x42(B)]  b1 b0 <binary-data>         # final chunk
//        ::= [x20-x2f] <binary-data>               # length less than 15
func encodeBinary(v []byte) (b []byte, err error) {
	if length := len(v); length <= 15 {
		b = append(b, 0x20+byte(length))
		b = append(b, v...)
		return b, err
	}
	
}

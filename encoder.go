package hessiancoder

import (
	"fmt"
	"log"
	"runtime"
	"time"
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

// EncodeBinary
// binary ::= [x62(b)]  b1 b0 <binary-data> binary  # non-final chunk
//        ::= [x42(B)]  b1 b0 <binary-data>         # final chunk
//        ::= [x20-x2f] <binary-data>               # length less than 15
func EncodeBinary(v []byte) (b []byte, err error) {
	if length := len(v); length <= 15 {
		b = append(b, 0x20+byte(length))
		b = append(b, v...)
		return b, err
	}

	offset := 0
	for length := len(v); length-offset > ChunkSize; {
		b = append(b, 'b')
		packB, err := packUInt16(ChunkSize)
		if err != nil {
			fmt.Println("can not packUInt16 for ", ChunkSize)
			return b, err
		}
		b = append(b, packB...)
		b = append(b, v[offset:ChunkSize]...)

		offset += ChunkSize
	}

	b = append(b, 'B')
	packB, err := packUInt16(uint16(len(v) - offset))
	if err != nil {
		fmt.Println("can not packUInt16 for ", len(v)-offset)
		return
	}
	b = append(b, packB...)
	b = append(b, v[offset:]...)
	return
}

// EncodeBool encode bool
// boolean ::= 'T'
//         ::= 'F'
func EncodeBool(f bool) (b []byte, err error) {
	if true == f {
		b = append(b, 'T')
		return
	}
	b = append(b, 'F')
	return
}

// EncodeDate encode date
// date    ::= 'd' b7 b6 b5 b4 b3 b2 b1 b0
func EncodeDate(date time.Time) (b []byte, err error) {
	tmpB := make([]byte, 0, 8)
	if tmpB, err = packInt64(date.UnixNano() / 1e6); err != nil {
		return
	}
	b = append(b, 'd')
	b = append(b, tmpB...)
	return
}

// EncodeFloat64 encode float64
// double    ::= 'D' b7 b6 b5 b4 b3 b2 b1 b0
func EncodeFloat64(v float64) (b []byte, err error) {
	tmpB := make([]byte, 0, 8)
	if tmpB, err = packFloat64(v); err != nil {
		return
	}
	b = append(b, 'D')
	b = append(b, tmpB...)
	return
}

// EncodeInt32 encode int
// int    ::= 'I' b3 b2 b1 b0
func EncodeInt32(v int32) (b []byte, err error) {
	tmpB := make([]byte, 0, 4)
	if tmpB, err = packInt32(v); err != nil {
		return
	}

	b = append(b, 'I')
	b = append(b, tmpB...)
	return
}

// EncodeLong encode long
// long    ::= 'L' b7 b6 b5 b4 b3 b2 b1 0
func EncodeLong(v int64) (b []byte, err error) {
	tmpB := make([]byte, 0, 8)
	if tmpB, err = packInt64(v); err != nil {
		return
	}

	b = append(b, 'L')
	b = append(b, tmpB...)
	return
}

// EncodeNull encode null
// null    ::= 'N'
func EncodeNull(v interface{}) (b []byte, err error) {
	b = append(b, 'N')
	return
}

// EncodeString encode string
// string    ::= R(x52) b1 b0 <utf8-data> string  # non-final chunk
//           ::= S(x53) b1 b0 <utf8-data>         # string of length 0-65535
//           ::= [x00-x1f] <utf8-data>            # string of length 0-31
//           ::= [x30-x33] b0 <utf8-data>         # string of length 0-1023
func EncodeString(s string) (b []byte, err error) {
	length, strOffset, subLen := len(s), 0, 0
	byteS := []byte(s)

	for length > ChunkSize {
		subLen = ChunkSize

		tmpB, err := packUInt16(uint16(subLen))
		if err != nil {
			return b, err
		}
		b = append(b, 'R')
		b = append(b, tmpB...)
		b = append(byteS[strOffset : strOffset+subLen])

		length -= subLen
		strOffset += subLen
	}

	if length <= 31 {
		// short strings
		b = append(b, byte(length))
	} else if length <= 1023 {
		b = append(b, byte(48+(length>>8)))
		b = append(b, byte(length))
	} else {
		b = append(b, 'S')
		b = append(b, byte(length>>8))
		b = append(b, byte(length))
	}
	b = append(b, byteS[strOffset:]...)
	return
}

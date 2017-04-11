package hessiancoder

import (
	"testing"
)

func TestEncodeBytes(t *testing.T) {
	lessThan15 := make([]byte, 0, 20)
	lessThan15 = append(lessThan15, 'a', 'b', 'c', 'd', 'e')
	result, err := EncodeBinary(lessThan15)
	if err != nil {
		t.Error("encode binary failed, lessThan15:", lessThan15, " error:", err)
	}
	t.Log("result of lessThan15: ", result)

	lessThan64k := make([]byte, 0, 0x8000)
	for i := 0; i < 0x7999; i++ {
		lessThan64k = append(lessThan64k, 'a')
	}
	result, err = EncodeBinary(lessThan64k)
	if err != nil {
		t.Error("encode binary failed, lessThan64k:", lessThan64k, " error:", err)
	}
	t.Log("result of lessThan64k: ", result)

	moreThan64k := make([]byte, 0, 0x8001)
	for i := 0; i <= 0x8000; i++ {
		moreThan64k = append(moreThan64k, 'a')
	}
	result, err = EncodeBinary(moreThan64k)
	if err != nil {
		t.Error("encode binary failed, moreThan64k:", moreThan64k, " error:", err)
	}
	t.Log("result of lessThan64k: ", result)
}

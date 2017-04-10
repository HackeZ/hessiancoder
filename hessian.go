package hessiancoder

import (
	"bufio"
)

type Any interface{}

type HessianCoder struct {
	coder *bufio.Reader
	refs  []Any
}
package iteration

import (
	"bytes"
	"fmt"
)

func RepeatBuf(char string, iter int) string {
	var ret bytes.Buffer
	for i := 0; i < iter; i++ {
		fmt.Fprint(&ret, char)
	}
	return ret.String()

}

func Repeat(char string, iter int) string {
	var ret string
	for i := 0; i < iter; i++ {
		ret += char
	}
	return ret

}

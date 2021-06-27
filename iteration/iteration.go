package iteration

import (
	"bytes"
	"fmt"
)

func RepeatBuf(char string) string {
	var ret bytes.Buffer
	for i := 0; i < 5000; i++ {
		fmt.Fprint(&ret, char)
	}
	return ret.String()

}

func Repeat(char string) string {
	var ret string
	for i := 0; i < 5000; i++ {
		ret += char
	}
	return ret

}

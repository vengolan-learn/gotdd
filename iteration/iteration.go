package iteration

import (
	"bytes"
	"fmt"
)

//Repeats char iter number of times.
//Uses Bytes.Buffer. This is more eifficient than Repeat
func RepeatBuf(char string, iter int) string {
	var ret bytes.Buffer
	for i := 0; i < iter; i++ {
		fmt.Fprint(&ret, char)
	}
	return ret.String()

}

//Repeats "char",  "iter" times
//This uses string concatenation for each iteration
func Repeat(char string, iter int) string {
	var ret string
	for i := 0; i < iter; i++ {
		ret += char
	}
	return ret

}

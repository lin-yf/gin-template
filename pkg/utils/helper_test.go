package utils

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func Test_Sonic(t *testing.T) {
	x := "chars@arefun"

	i := strings.Index(x, "@a")
	fmt.Println("Index: ", i)
}

func Test_extract(t *testing.T) {
	s := "在最近几年里，我写了很多有关HTTPS和HTTP/2的文章，涵盖了证书申请"
	str := []rune(s)
	fmt.Println(strings.IndexRune(s, '我'))
	b := str[0:11]
	fmt.Println(string(b))
}

func Test_substring(t *testing.T) {
	input_string := "hello world 您好你好"
	byte_index := strings.IndexRune(input_string, '你')
	fmt.Println(byte_index)
	fmt.Println(input_string[byte_index-10 : byte_index+6])
	fmt.Println(utf8.RuneCountInString(input_string[:byte_index]))
}

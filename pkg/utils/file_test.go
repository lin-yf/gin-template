package utils

import (
	"fmt"
	"testing"
)

func Test_File(m *testing.T) {
	fmt.Println("test")
	err := DelFile()
	fmt.Println("test", err)
}

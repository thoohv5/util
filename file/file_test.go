package file

import (
	"fmt"
	"testing"
)

func TestWriteFile(t *testing.T) {
	str := "111"
	fmt.Println(Write("./doc/1.txt", str))
	fmt.Println(Write("./doc/1.txt", str))
	fmt.Println(Write("./doc/1.txt", str))
	fmt.Println(Read("./doc/1.txt"))
}

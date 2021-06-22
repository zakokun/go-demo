package main

import (
	"bytes"
	"fmt"
)

func main() {
	var bf bytes.Buffer
	bf.WriteString("a")
	bf.WriteString("b")
	bf.WriteString("c")
	bf.Truncate(bf.Len()-1)
	fmt.Println(bf.String())
}

package main

import (
        "fmt"
        "unicode"
)

func main() {
        fmt.Println(unicode.Is(unicode.Han,[]rune("わはは")[0]))
}

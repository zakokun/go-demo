package main

import "github.com/davecgh/go-spew/spew"

func main() {
	ls := []int64{1, 12, 5, 24, 5324, 6, 8, 4, 2, 5, 8, 9, 5}
	for i := 0; i < len(ls); i++ {
		for j := i + 1; j < len(ls); j++ {
			if ls[i] < ls[j] {
				tp := ls[i]
				ls[i] = ls[j]
				ls[j] = tp
			}
		}
	}

	spew.Dump(ls)
}

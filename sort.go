package main

import "github.com/davecgh/go-spew/spew"

//func main() {
//	ls := []int64{1, 12, 5, 24, 5324, 6, 8, 4, 2, 5, 8, 9, 5}
//	for i := 0; i < len(ls); i++ {
//		for j := i + 1; j < len(ls); j++ {
//			if ls[i] < ls[j] {
//				tp := ls[i]
//				ls[i] = ls[j]
//				ls[j] = tp
//			}
//		}
//	}
//	spew.Dump(ls)
//}

func main() {
	ls := []int64{123, 4, 2, 44, 23, 34, 235, 346, 23, 1212, 3123, 123, 1, 3}
	//selectSort(ls)
	//spew.Dump(ls)

	bubbleSort(ls)
	spew.Dump(ls)
}

func bubbleSort(ls []int64) {
	for i := 0; i < len(ls); i++ {
		swap := false
		for j := i + 1; j < len(ls); j++ {
			if ls[i] < ls[j] {
				tp := ls[i]
				ls[i] = ls[j]
				ls[j] = tp
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

func selectSort(ls []int64) {
	for i := 0; i < len(ls); i++ {
		flag := i
		for j := i + 1; j < len(ls); j++ {
			if ls[flag] < ls[j] {
				flag = j
			}
		}
		if flag != i {
			tp := ls[i]
			ls[i] = ls[flag]
			ls[flag] = tp
		}
	}
}

func insertSort(ls []int64) {
	for i := 1; i < len(ls); i++ {
		for j := i + 1; j > 0; j-- {

		}
	}
}

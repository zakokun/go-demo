package main

import "github.com/davecgh/go-spew/spew"

func main() {
	mp := make(map[int64]int64)

	mp[1] = 11
	mp[2] = 22
	mp[3] = 33
	mp[4] = 44

	for k := range mp {
		delete(mp, k)
	}
	spew.Dump(mp)
}

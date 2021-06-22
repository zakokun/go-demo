package main

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {
	region, err := ip2region.New("data/ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	ip, _ := region.MemorySearch("27.115.6.196")
	fmt.Println(ip.Region)

}
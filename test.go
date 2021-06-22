package main

import "fmt"
import "regexp"

func main() {
	rex := "/glados/admin/[^(nologin)].*"
	reg := regexp.MustCompile(rex)
	fmt.Println(reg.FindStringIndex("/glados/admin/nogin/aabb/dd"))
}
package main

import "fmt"

func main() {
	set := make(map[string]bool)

	set["red"] = true
	set["green"] = true
	set["blue"] = false

	fmt.Println(set)
}

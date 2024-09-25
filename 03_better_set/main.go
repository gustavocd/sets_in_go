package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	colors := mapset.NewSet[string]()
	colors.Add("red")
	colors.Add("green")
	colors.Add("blue")
	colors.Add("orange")
	colors.Add("purple")

	rgb := mapset.NewSet[string]()
	rgb.Add("red")
	rgb.Add("green")
	rgb.Add("blue")

	both := colors.Intersect(rgb)

	fmt.Println(both)
}

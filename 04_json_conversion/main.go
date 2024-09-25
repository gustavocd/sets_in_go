package main

import (
	"encoding/json"
	"fmt"
	"log"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	colors := mapset.NewSet[string]()
	colors.Add("red")
	colors.Add("green")
	colors.Add("blue")

	data, err := json.Marshal(colors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

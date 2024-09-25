package main

import "fmt"

func main() {
	set := map[string]struct{}{
		"red":   {},
		"green": {},
	}
	fmt.Println("Initial set:", set)

	// Adding elements to the set
	set["black"] = struct{}{}
	set["blue"] = struct{}{}
	fmt.Println("After adding:", set)

	// Removing elements of a set
	delete(set, "red")
	delete(set, "green")
	fmt.Println("After removing:", set)

	// Checking for existing element
	_, ok := set["black"]
	fmt.Printf("is black in the set? %v\n", ok)

	val, ok := set["orange"]
	fmt.Printf("is orange in the set? %v\n", ok)
	fmt.Printf("is orange an empty struct? %v\n", val == struct{}{})

	// Size of the set (also known as "cardinality")
	size := len(set)
	fmt.Println("The set size is:", size)

	// Listing elements
	for color := range set {
		fmt.Println(color)
	}

	// What about more functionality? things like:
	// 1. Difference between two sets.
	// 2. Intersection of two sets.
	// 3. Checking for subset/superset.
}

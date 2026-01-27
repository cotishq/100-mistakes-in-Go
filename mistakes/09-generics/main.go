package main

import (
	"fmt"
)

func main() {
	fmt.Println("Before generics")
	m1 := map[string]int{"one": 1, "two": 2}
	fmt.Println(getStringIntkeys(m1))

	fmt.Println("After generics")
	m2 := map[int]string{1: "one", 2: "two"}
	fmt.Println(getKeys(m1))
	fmt.Println(getKeys(m2))
}



func getStringIntkeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys

}

//if we want map[int]string keys, we need another function
func getIntStringKeys(m map[int]string) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	return keys

}

// Correct use of generics
// K must be comparable because map keys must be comparable
func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}



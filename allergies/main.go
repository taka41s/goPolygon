package main

import (
	"fmt"
	"exercism/allergies"
)

func main() {
	score := 34

	item := "cats"
	if allergies.IsAllergicTo(score, item) {
		fmt.Printf("Allergic to %s\n", item)
	} else {
		fmt.Printf("Not allergic to %s\n", item)
	}

	allergies := allergies.ListAllergies(score)
	fmt.Println("Allergic to:", allergies)
}

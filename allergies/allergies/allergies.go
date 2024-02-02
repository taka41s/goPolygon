package allergies

import (
	"fmt"
)

var allergens = map[string]int{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func IsAllergicTo(score int, item string) bool {
	itemScore, exists := allergens[item]
	fmt.Printf("Binary representation of score: %b\n", score)
	if !exists {
		fmt.Printf("%s is not a recognized allergen.\n", item)
		return false
	}
	return score&itemScore == itemScore
}

func ListAllergies(score int) []string {
	var allergies []string
	for allergen, value := range allergens {			// 34 = 100010
		if score&value == value { 						//100000 == value //000010 == value
 			allergies = append(allergies, allergen)		
		}
	}
	return allergies
}

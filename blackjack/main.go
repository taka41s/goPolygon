package main

import "fmt"

func ParseCard(card string) int {
	cardValues := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"other": 0,
	}

	return cardValues[card]
}

func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	fmt.Println(sum)
	switch {
	case card1 == card2: 
		return "P"
	case sum == 21 && (dealerCard != "ace" && dealerCard != "king" && dealerCard != "queen" && dealerCard != "ten"):
		return "W"
	case sum >= 17 && sum <= 21 :
		return "S" 
	case sum >= 12 && sum <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H" 
		}
		return "S" 
	default:
		return "H" 
	}
	
}

func main() {
	fmt.Println(FirstTurn("ace", "ace", "jack"))
	fmt.Println(FirstTurn("ace", "king", "ace")) 
	fmt.Println(FirstTurn("five", "queen", "ace"))
}
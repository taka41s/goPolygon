package main
import (
	"fmt"
	"gobase/math"
)

func main(){
	calc(3, 0, "div")
}

func calc(num1 int, num2 int, operation string) {
	switch  operation {
		case "sum":
			fmt.Println(math.Sum(num1, num2))
		case "sub":
			fmt.Println(math.Sub(num1, num2))
		case "div":
			fmt.Println(math.Div(num1, num2))
		case "multi":
			fmt.Println(math.Multi(num1, num2))
	}
}


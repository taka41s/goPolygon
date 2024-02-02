package main

import ("fmt" 
		"errors")

type ErrNegativeSqrt float64

func Abs(n float64) float64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %.2f\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		z := x

		for Abs(x-z*z) > 0.001 {

			z -= (z*z - x) / (2 * z)

		}

		return z, nil

	} else {
		return 0, ErrNegativeSqrt(x)

	}
}

func Divide(number1 float64, number2 float64) (float64, error) {
	if number2 == 0 {
		return 0, errors.New("division by zero")
	}

	result := number1 / number2
	return result, nil
}

func main() {
	division, err := Divide(42.4, 55.5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(division)
	}
}
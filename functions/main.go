package main

import (
	"fmt"
)
func main() {
	var printValue string = "hi"
	printMe(printValue)

	var numerator int = 11
	var denominator = 2
	var result, remainder int = intDivision(numerator, denominator)
	fmt.Printf("The result of the integer is %v with remainder %w", result, remainder)
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var err error
	if denominator==0{
		return err, err
	}
	var result int = numerator/denominator
	var remainder int =	numerator%denominator
	return result, remainder
}

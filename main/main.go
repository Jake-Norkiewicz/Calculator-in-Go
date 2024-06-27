package main

import (
	"fmt"
	"mymodule/mypackage"
)

func main() {
	var x, y float64
	
	fmt.Println("Please enter two numbers: ")
	fmt.Scan(&x)
	fmt.Scan(&y)

	sum := mypackage.Add(x, y)
	fmt.Printf("Sum: %f", sum)
}
package main

import (
	"fmt"
)


func main() {
fmt.Println("Enter two numbers to add")
var a,b,c int
    _, err := fmt.Scanf("%d %d", &a,&b) // space is important for scanning 2 numbers
	_ =  err
	c =  a + b
	fmt.Println("Sum of entered numbers = ",c)

}
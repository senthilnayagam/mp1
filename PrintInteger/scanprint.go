package main

import (
	"fmt"
)


func main() {
fmt.Println("Enter an integer")
var a int
    _, err := fmt.Scanf("%d", &a)
	_ =  err
	fmt.Println("Integer that you have entered is ",a)

}
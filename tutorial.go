package main

import	"fmt"
import t "time"



func main(){
	//print my first statement
	fmt.Println("Hy my first golang program!")
	//using go's import statement
	fmt.Println(t.Now())
	//go formatting
	fmt.Printf("Number: %T,%v", 100, 100)
	//binary formatting
	fmt.Printf("Binary formatting: %b", 3435)	

}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the year you were born: ")
	scanner.Scan()

	//the below line converts a string to int, 10 is the base of the number its converting to and its 64 bits
	//the _ stores the error and since we are not interested we use _
	input, _ := strconv.ParseInt(scanner.Text(), 10 , 64)

	fmt.Printf("You will be %d years old at the end of 2021", 2021-input)
}
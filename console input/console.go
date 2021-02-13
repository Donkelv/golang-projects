package main

import (
	"bufio"
	"fmt"
	"os"
	
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your name: ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("You typed: %q", input)
}
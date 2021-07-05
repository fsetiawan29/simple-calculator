package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter first value: ")
		scanner.Scan()
		firstValue := scanner.Text()
		fmt.Println(firstValue)

		fmt.Printf("Enter second value: ")
		scanner.Scan()
		secondValue := scanner.Text()
		fmt.Println(secondValue)

		break
	}
}

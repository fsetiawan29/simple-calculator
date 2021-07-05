package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fsetiawan29/simple-calculator/module"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter first value: ")
		scanner.Scan()
		firstValue := scanner.Text()
		firstValueInt, err := strconv.Atoi(firstValue)
		if err != nil {
			fmt.Println("Please input an integer")
			break
		}

		fmt.Printf("Enter second value: ")
		scanner.Scan()
		secondValue := scanner.Text()
		secondValueInt, err := strconv.Atoi(secondValue)
		if err != nil {
			fmt.Println("Please input an integer")
			break
		}

		result := module.Add(firstValueInt, secondValueInt)

		fmt.Printf("The result is %d\n", result)

		break
	}
}

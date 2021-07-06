package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fsetiawan29/simple-calculator/module"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("List Operation")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Printf("Select Operation: ")
		scanner.Scan()
		operation := scanner.Text()

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

		var result int
		if operation == "1" {
			result = module.Add(firstValueInt, secondValueInt)
			if result < 0 {
				log.Println(result)
			}
		} else if operation == "2" {
			result = module.Subtract(firstValueInt, secondValueInt)
			if result < 0 {
				log.Println(result)
			}
		} else if operation == "3" {
			result = module.Multiply(firstValueInt, secondValueInt)
			if result < 0 {
				log.Println(result)
			}
		} else if operation == "4" {
			result = module.Divide(firstValueInt, secondValueInt)
			if result < 0 {
				log.Println(result)
			}
		} else {
			fmt.Println("Please select correct operation")
			break
		}

		fmt.Printf("The result is %d\n", result)

		break
	}
}

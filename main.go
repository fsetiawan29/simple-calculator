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
		firstValueInt, _ := strconv.Atoi(firstValue)

		fmt.Printf("Enter second value: ")
		scanner.Scan()
		secondValue := scanner.Text()
		secondValueInt, _ := strconv.Atoi(secondValue)

		result := module.Add(firstValueInt, secondValueInt)

		fmt.Printf("The result is %d\n", result)

		break
	}
}

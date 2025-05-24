package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetUserInput(hint string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(hint)
	scanner.Scan()
	input := scanner.Text()

	if input == "" {
		fmt.Println("Can not be empty.")
		return GetUserInput(hint)
	}

	return input
}

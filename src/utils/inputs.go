package utils

import (
	"fmt"
	"os"
)

func InputAsString(file string) string {
	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	i := string(input)
	return i
}

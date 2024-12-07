package utils

import (
	"log"
	"os"
)

func InputAsString(file string) string {
	input, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	i := string(input)
	return i
}

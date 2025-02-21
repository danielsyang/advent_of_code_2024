package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Read_file(file_path string) *bufio.Scanner {
	file, err := os.Open(file_path)

	if err != nil {
		fmt.Println("Can't read filepath:", file_path)
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	return scanner
}

func Abs_int(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

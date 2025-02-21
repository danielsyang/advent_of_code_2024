package day_one

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func parse_file(file_path string) (first, last []int) {
	var first_list []int
	var last_list []int

	file, err := os.Open(file_path)

	if err != nil {
		fmt.Println("Can't read filepath:", file_path)
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		size := len(line) - 1

		first, err := strconv.Atoi(string(line[0]))
		if err != nil {
			fmt.Println("Can't format values first: ", first)
			panic(err)
		}

		last, err := strconv.Atoi(string(line[size]))
		if err != nil {
			fmt.Println("Can't format values last: ", last)
			panic(err)
		}

		first_list = append(first_list, first)
		last_list = append(last_list, last)
	}

	return first_list, last_list
}

func Day_one(file_path string) {
	first_list, second_list := parse_file(file_path)
	total := 0

	slices.Sort(first_list)
	slices.Sort(second_list)

	fmt.Println(first_list)
	fmt.Println(second_list)

	for index, value := range first_list {
		if value > second_list[index] {
			total += value - second_list[index]
		} else {
			total += second_list[index] - value
		}
	}

	fmt.Println("What is the total distance between your lists?", total)

}

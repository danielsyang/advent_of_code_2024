package day_one

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/danielsyang/go_advent_of_code/utils"
)

func parse_file(file_path string) (first, last []int) {
	var first_list []int
	var last_list []int
	scanner := utils.Read_file(file_path)

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

	for index, value := range first_list {
		total += utils.Abs_int(value - second_list[index])
	}

	fmt.Println("What is the total distance between your lists?", total)

}

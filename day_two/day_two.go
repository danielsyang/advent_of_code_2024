package day_two

import (
	"fmt"
	"strconv"

	"github.com/danielsyang/go_advent_of_code/utils"
)

func parse_file(file_path string) [][]int {

	var report [][]int
	scanner := utils.Read_file(file_path)

	for scanner.Scan() {
		line := scanner.Text()

		var levels []int

		for _, r := range line {
			number, err := strconv.Atoi(string(r))

			if err == nil {
				levels = append(levels, number)
			}

		}
		report = append(report, levels)
	}

	return report

}

func Day_two(file_path string) {

	report := parse_file(file_path)
	var total = 0

	for _, levels := range report {

		var is_increasing bool
		is_valid := true

		for i := range len(levels) - 1 {
			j := i + 1

			if !is_valid {
				break
			}

			if i == 0 {
				if levels[i] > levels[j] {
					is_increasing = true
				} else {
					is_increasing = false
				}
			}

			if levels[i] > levels[j] && !is_increasing || levels[i] < levels[j] && is_increasing || levels[i] == levels[j] {
				is_valid = false
			} else if utils.Abs_int(levels[i]-levels[j]) <= 3 {
				is_valid = true
			} else {
				is_valid = false
			}

		}

		if is_valid {
			total++
		}

	}

	fmt.Println("How many reports are safe?", total)

}

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

	reports := parse_file(file_path)

	fmt.Println(reports)
}

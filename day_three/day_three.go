package day_three

import (
	"fmt"

	"github.com/danielsyang/go_advent_of_code/utils"
)

func parse_file(file_path string) [][]int {

	var report [][]int
	scanner := utils.Read_file(file_path)

	for scanner.Scan() {
		line := scanner.Text()

		for i := range len(line) {

			test := string(line[i])
			fmt.Println(test)
		}

	}

	return report

}

func Day_three(file_path string) {
	parse_file(file_path)
}

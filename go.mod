module github.com/danielsyang/go_advent_of_code

go 1.24.0

replace github.com/danielsyang/go_advent_of_code/day_one => ./day_one

replace github.com/danielsyang/go_advent_of_code/day_two => ./day_two

replace github.com/danielsyang/go_advent_of_code/day_three => ./day_three

replace github.com/danielsyang/go_advent_of_code/utils => ./utils

require (
	github.com/danielsyang/go_advent_of_code/day_one v0.0.0-00010101000000-000000000000
	github.com/danielsyang/go_advent_of_code/day_three v0.0.0-00010101000000-000000000000
	github.com/danielsyang/go_advent_of_code/day_two v0.0.0-00010101000000-000000000000
)

require github.com/danielsyang/go_advent_of_code/utils v0.0.0-00010101000000-000000000000 // indirect

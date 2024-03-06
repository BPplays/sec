package main

import (
	"strings"
	"time"
)

var date_sep = map[rune]bool{
	'/': true,
	'-': true,
	'.': true,
}

var date_time_sep = map[rune]bool{
	' ': true,
	'T': true,
}

var time_sep = map[rune]bool{
	':': true,
}

var builders = map[string]strings.Builder{
	"year": strings.Builder{},
	"mon": strings.Builder{},
	"day": strings.Builder{},
	"time": strings.Builder{},
}


func parse_date(in string) *time.Time {

	// var year strings.Builder
	// var mon strings.Builder
	// var day strings.Builder
	// var time strings.Builder
	// fmt.Println(input)
	// runel := []rune(input)

	for _, builder := range builders {
		for _, i := range in {
			if date_sep[i] {
				continue
			} else {
				builder.WriteRune(i)
			}
		}

	}

	// return strconv.ParseInt(sb.String(), 10, 64)
}
package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func exerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
	"hr": strings.Builder{},
	"min": strings.Builder{},
	"sec": strings.Builder{},
}

var builders_sep = map[string]map[rune]bool{
	"year": date_sep,
	"mon": date_sep,
	"day": date_time_sep,
	"hr": time_sep,
	"min": time_sep,
	"sec": time_sep,
}


func parse_date(in string) time.Time {
	var out_time time.Time

	for builder_name, builder := range builders {
		for _, i := range in {
			if builders_sep[builder_name][i] {
				continue
			} else {
				builder.WriteRune(i)
				// builder.String()
			}
		}

	}

	year, err := strconv.Atoi(builders["year"].String())
	exerr(err)
	mon, err := strconv.Atoi(builders["mon"].String())
	exerr(err)
	day, err := strconv.Atoi(builders["day"].String())
	exerr(err)
	hr, err := strconv.Atoi(builders["hr"].String())
	exerr(err)
	min, err := strconv.Atoi(builders["min"].String())
	exerr(err)
	sec, err := strconv.Atoi(builders["sec"].String())
	exerr(err)
	// out_time.AddDate()
	out_time = time.Date(year, time.Month(mon), day, hr, min, sec, 0, 0, time.UTC)
	if err != nil {
		log.Fatal(err)
	}
	// return strconv.ParseInt(sb.String(), 10, 64)
	return out_time
}
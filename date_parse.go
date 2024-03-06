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

	var b strings.Builder

	b = builders["year"]
	year, err := strconv.Atoi(b.String())
	exerr(err)

	b = builders["mon"]
	mon, err := strconv.Atoi(b.String())
	exerr(err)

	b = builders["day"]
	day, err := strconv.Atoi(b.String())
	exerr(err)

	b = builders["hr"]
	hr, err := strconv.Atoi(b.String())
	exerr(err)

	b = builders["min"]
	min, err := strconv.Atoi(b.String())
	exerr(err)

	b = builders["sec"]
	sec, err := strconv.Atoi(b.String())
	exerr(err)
	// out_time.AddDate()
	out_time = time.Date(year, time.Month(mon), day, hr, min, sec, 0, time.UTC)
	if err != nil {
		log.Fatal(err)
	}
	// return strconv.ParseInt(sb.String(), 10, 64)
	return out_time
}
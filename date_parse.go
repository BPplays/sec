package main

import (
	"fmt"
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

var builders = map[string]*strings.Builder{
	"year": new(strings.Builder),
	"mon":  new(strings.Builder),
	"day":  new(strings.Builder),
	"hr":   new(strings.Builder),
	"min":  new(strings.Builder),
	"sec":  new(strings.Builder),
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

	runel := []rune(in)

	for builder_name, builder := range builders {
		for _, i := range runel {
			if builders_sep[builder_name][i] {
				break
			} else {
				builders[builder_name].WriteRune(i)
				fmt.Println(builder_name, " ", builder.String())
			}
		}
		runel = runel[1:]

	}

	// var b strings.Builder

	// b = builders["year"]
	year, err := strconv.Atoi(builders["year"].String())
	if err != nil {
		log.Fatal("yr", err)
	}

	// b = builders["mon"]
	mon, err := strconv.Atoi(builders["mon"].String())
	exerr(err)

	// b = builders["day"]
	day, err := strconv.Atoi(builders["day"].String())
	exerr(err)

	// b = builders["hr"]
	hr, err := strconv.Atoi(builders["hr"].String())
	exerr(err)

	// b = builders["min"]
	min, err := strconv.Atoi(builders["min"].String())
	exerr(err)

	// b = builders["sec"]
	sec, err := strconv.Atoi(builders["sec"].String())
	exerr(err)
	// out_time.AddDate()
	out_time = time.Date(year, time.Month(mon), day, hr, min, sec, 0, time.UTC)
	if err != nil {
		log.Fatal(err)
	}
	// return strconv.ParseInt(sb.String(), 10, 64)
	return out_time
}
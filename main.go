package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	// "github.com/BPplays/dateparse"
	"github.com/araddon/dateparse"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)


type Prefix struct {
	Symbol    string
	Base10    float64
	FullName  string
	Adoption  int
}

var AllPrefixes = map[string]Prefix{
	"quetta": {Symbol: "Q", Base10: math.Pow(10, 30), FullName: "quetta", Adoption: 2022},
	"ronna":  {Symbol: "R", Base10: math.Pow(10, 27), FullName: "ronna", Adoption: 2022},
	"yotta":  {Symbol: "Y", Base10: math.Pow(10, 24), FullName: "yotta", Adoption: 1991},
	"zetta":  {Symbol: "Z", Base10: math.Pow(10, 21), FullName: "zetta", Adoption: 1991},
	"exa":    {Symbol: "E", Base10: math.Pow(10, 18), FullName: "exa", Adoption: 1975},
	"peta":   {Symbol: "P", Base10: math.Pow(10, 15), FullName: "peta", Adoption: 1975},
	"tera":   {Symbol: "T", Base10: math.Pow(10, 12), FullName: "tera", Adoption: 1960},
	"giga":   {Symbol: "G", Base10: math.Pow(10, 9), FullName: "giga", Adoption: 1960},
	"mega":   {Symbol: "M", Base10: math.Pow(10, 6), FullName: "mega", Adoption: 1873},
	"kilo":   {Symbol: "k", Base10: math.Pow(10, 3), FullName: "kilo", Adoption: 1795},
	"hecto":  {Symbol: "h", Base10: math.Pow(10, 2), FullName: "hecto", Adoption: 1795},
	"deca":   {Symbol: "da", Base10: math.Pow(10, 1), FullName: "deca", Adoption: 1795},
	"none":   {Symbol: "", Base10: math.Pow(10, 0), FullName: "none", Adoption: 1795},
	"deci":   {Symbol: "d", Base10: math.Pow(10, -1), FullName: "deci", Adoption: 1795},
	"centi":  {Symbol: "c", Base10: math.Pow(10, -2), FullName: "centi", Adoption: 1795},
	"milli":  {Symbol: "m", Base10: math.Pow(10, -3), FullName: "milli", Adoption: 1795},
	"micro":  {Symbol: "µ", Base10: math.Pow(10, -6), FullName: "micro", Adoption: 1873},
	"nano":   {Symbol: "n", Base10: math.Pow(10, -9), FullName: "nano", Adoption: 1960},
	"pico":   {Symbol: "p", Base10: math.Pow(10, -12), FullName: "pico", Adoption: 1960},
	"femto":  {Symbol: "f", Base10: math.Pow(10, -15), FullName: "femto", Adoption: 1964},
	"atto":   {Symbol: "a", Base10: math.Pow(10, -18), FullName: "atto", Adoption: 1964},
	"zepto":  {Symbol: "z", Base10: math.Pow(10, -21), FullName: "zepto", Adoption: 1991},
	"yocto":  {Symbol: "y", Base10: math.Pow(10, -24), FullName: "yocto", Adoption: 1991},
	"ronto":  {Symbol: "r", Base10: math.Pow(10, -27), FullName: "ronto", Adoption: 2022},
	"quecto": {Symbol: "q", Base10: math.Pow(10, -30), FullName: "quecto", Adoption: 2022},
}

var common_prefixes = map[string]Prefix{
	"quetta": {Symbol: "Q", Base10: math.Pow(10, 30), FullName: "quetta", Adoption: 2022},
	"ronna":  {Symbol: "R", Base10: math.Pow(10, 27), FullName: "ronna", Adoption: 2022},
	"yotta":  {Symbol: "Y", Base10: math.Pow(10, 24), FullName: "yotta", Adoption: 1991},
	"zetta":  {Symbol: "Z", Base10: math.Pow(10, 21), FullName: "zetta", Adoption: 1991},
	"exa":    {Symbol: "E", Base10: math.Pow(10, 18), FullName: "exa", Adoption: 1975},
	"peta":   {Symbol: "P", Base10: math.Pow(10, 15), FullName: "peta", Adoption: 1975},
	"tera":   {Symbol: "T", Base10: math.Pow(10, 12), FullName: "tera", Adoption: 1960},
	"giga":   {Symbol: "G", Base10: math.Pow(10, 9), FullName: "giga", Adoption: 1960},
	"mega":   {Symbol: "M", Base10: math.Pow(10, 6), FullName: "mega", Adoption: 1873},
	"kilo":   {Symbol: "k", Base10: math.Pow(10, 3), FullName: "kilo", Adoption: 1795},
	"none":   {Symbol: "", Base10: math.Pow(10, 0), FullName: "none", Adoption: 1795},
	"milli":  {Symbol: "m", Base10: math.Pow(10, -3), FullName: "milli", Adoption: 1795},
	"micro":  {Symbol: "µ", Base10: math.Pow(10, -6), FullName: "micro", Adoption: 1873},
	"nano":   {Symbol: "n", Base10: math.Pow(10, -9), FullName: "nano", Adoption: 1960},
	"pico":   {Symbol: "p", Base10: math.Pow(10, -12), FullName: "pico", Adoption: 1960},
	"femto":  {Symbol: "f", Base10: math.Pow(10, -15), FullName: "femto", Adoption: 1964},
	"atto":   {Symbol: "a", Base10: math.Pow(10, -18), FullName: "atto", Adoption: 1964},
	"zepto":  {Symbol: "z", Base10: math.Pow(10, -21), FullName: "zepto", Adoption: 1991},
	"yocto":  {Symbol: "y", Base10: math.Pow(10, -24), FullName: "yocto", Adoption: 1991},
	"ronto":  {Symbol: "r", Base10: math.Pow(10, -27), FullName: "ronto", Adoption: 2022},
	"quecto": {Symbol: "q", Base10: math.Pow(10, -30), FullName: "quecto", Adoption: 2022},
}




func fmt_epoch_to_prefixsec(utime int64, prefixes map[string]Prefix, break_prefix string, mul *float64) string {
	var output strings.Builder

	var fl_time float64

	if mul != nil {
		fl_time = float64(utime) * *(mul)
	} else {
		fl_time = float64(utime)
	}
	
	var fl_round_time float64


	keys := make([]string, 0, len(prefixes))
	for key := range prefixes {
		keys = append(keys, key)
	}

	// Sort the keys in descending order
	sort.Slice(keys, func(i, j int) bool {
		return prefixes[keys[i]].Base10 > prefixes[keys[j]].Base10
	})

	// Iterate over the sorted keys
	for _, key := range keys {
		value := prefixes[key]
		if key == break_prefix {
			break
		}


		if fl_time / value.Base10 >= 1 {
			fl_round_time = math.Floor(fl_time / value.Base10)
			output.WriteString(fmt.Sprintf("%v%v",fl_round_time, value.Symbol+"s"))
			fl_time = fl_time - (fl_round_time * value.Base10)
			output.WriteString(" ")
		}

	}

	return removeSingleTrailingSpace(output.String())
}


func removeSingleTrailingSpace(input string) string {
	// Check if the input string has a single trailing space
	if strings.HasSuffix(input, " ") {
		// If yes, remove the last character
		return input[:len(input)-1]
	}
	// If no trailing space, return the input string as is
	return input
}


func findAndParseNumber(input string) (int64, error) {
	// Define a regular expression to match the first group of numbers
	re := regexp.MustCompile(`(\d+)`)

	// Find the first match
	match := re.FindStringSubmatch(input)

	// if len(match) < 2 {
	// 	return 0, fmt.Errorf("No numbers found in the input string")
	// }

	// Parse the matched numbers as int64
	number, err := strconv.ParseInt(match[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}


func parse_prefix_sec(input string) int64 {
	var total int64 = 0

	split := strings.SplitAfter(input, "s")





	var rune_list []rune

	for _, i := range split {
		num, err := findAndParseNumber(i)
		if err != nil {
			continue
		}

		rune_list = []rune(i)
		prefix := rune_list[len(rune_list)-2]
		if prefix == ' ' || unicode.IsDigit(prefix) {
			total += num
			continue
		}

		for _, value := range AllPrefixes {
			// Check if the Symbol matches the symbolToFind
			if value.Symbol == string(prefix) {
				// Symbol found, do something
				// fmt.Printf("Symbol %s found for prefix %s\n", prefix, key)
				total += num * int64(value.Base10)
			}
		}
	}

	return total

}








func main() {

	var epochTime int64

	var utime *int64
	var millisecflag bool
	var microsecflag bool
	var nanosecflag bool

	var baresecflag bool

	var date string

	var prefix_second string

	var date_out bool


	var break_prefix string = "milli"

	// Set up the command line flags
	pflag.Int64P("int_second", "i", 0, "int_second input")
	pflag.StringVarP(&prefix_second, "prefix_second", "p", "", "prefix_second input")


	pflag.BoolVarP(&millisecflag, "milli", "m", false, "milliseconds")
	pflag.BoolVarP(&microsecflag, "micro", "6", false, "microseconds (6 is for 10^-6 what micro stands for)")
	pflag.BoolVarP(&nanosecflag, "nano", "n", false, "nanoseconds")

	pflag.BoolVarP(&baresecflag, "bare", "b", false, "bareseconds format")

	pflag.StringVarP(&date, "date", "d", "", "date format")

	pflag.BoolVarP(&date_out, "date_out", "o", false, "date output")


	pflag.Parse()

	// Bind the viper configuration to the command line flags
	viper.BindPFlags(pflag.CommandLine)

	// Get the utime value from the configuration
	if viper.IsSet("int_second") {
		utimeValue := viper.GetInt64("int_second")
		utime = &utimeValue
	} else {
		utime = nil
	}


	millisec := viper.GetBool("milli")
	microsec := viper.GetBool("micro")
	nanosec := viper.GetBool("nano")

	baresec := viper.GetBool("bare")

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the Unix epoch time in seconds

	var mul_val float64 = 1
	
	var mul *float64

	mul = &mul_val

	if viper.IsSet("prefix_second") {
		utimeValue := viper.GetString("prefix_second")
		utimeValue_parse := parse_prefix_sec(utimeValue)
		utime = &utimeValue_parse
	}



	if millisec {
		epochTime = currentTime.UnixMilli()
		break_prefix = "micro"
		*(mul) = math.Pow(10, -3)
	} else if microsec {
		epochTime = currentTime.UnixMicro()
		break_prefix = "nano"
		*(mul) = math.Pow(10, -6)
	} else if nanosec {
		epochTime = currentTime.UnixNano()
		break_prefix = "pico"
		*(mul) = math.Pow(10, -9)
	} else {
		epochTime = currentTime.Unix()
		mul = nil
	}


	if utime == nil {
		// fmt.Println("utime is not assigned. Using default value.")
		utime = &epochTime
	}

	// time.

	if baresec{
		fmt.Println((*utime))
	} else if date != "" {
		// date, err := time.Parse(customLayout, date)
		// format := "%Y/%m/%d %H:%M"
		// date_p, err := strftime.Parse(date, format)
		parsed_date, err := dateparse.ParseStrict(date)
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println(parsed_date)
		fmt.Println(parsed_date.Unix())
	} else if date_out {
		var date_out time.Time

		var format string

		if millisec {
			date_out = time.UnixMilli((*utime))
			format = "2006-01-02 15:04:05.000"
		} else if microsec {
			log.Fatal("error cant use microsec and date output")
			date_out = time.UnixMicro((*utime))
		} else if nanosec {
			log.Fatal("error cant use nanosec and date output")
			// date_out = time.UnixNano(utime)
		} else {
			date_out = time.Unix((*utime), 0)
			format = "2006-01-02 15:04:05"
		}

		fmt.Printf("local: %v\nUTC: %v\n",date_out.Format(format) ,date_out.UTC().Format(format))
	} else {
		fmt.Println(fmt_epoch_to_prefixsec((*utime), common_prefixes, break_prefix, mul))
	}

}

package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	// "github.com/BPplays/dateparse"

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

var leading_zero bool


func fmt_epoch_to_prefixsec(utime int64, prefixesp *map[string]Prefix, break_prefix string, mul *float64) string {
	var output strings.Builder

	var fl_time float64

	prefixes := *(prefixesp)

	if mul != nil {
		fl_time = float64(utime) * *(mul)
	} else {
		fl_time = float64(utime)
	}

	if round_power != 0 {
		// fl_time = math.Floor(fl_time / math.Pow10(int(round_power))) * math.Pow10(int(round_power))
		fl_time = fl_time - (math.Mod(fl_time, float64(math.Pow10(int(round_power)))))
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



	var value Prefix
	var next_value Prefix
	var powerDifference float64

	var first_non0 bool = false
	// Iterate over the sorted keys
	for i, key := range keys {
		value = prefixes[key]

		if i+1 < len(prefixes) {
			next_value = prefixes[keys[i+1]]

			powerDifference = math.Log10(value.Base10) - math.Log10(next_value.Base10)
		} else {
			powerDifference = 3
		}




		if fl_time / value.Base10 >= 1 || (show_all_values && first_non0) || show_all_values_super {
			fl_round_time = math.Floor(fl_time / value.Base10)
			// fl_round_time = math.Floor(fl_round_time / math.Pow10(int(round_power))) * math.Pow10(int(round_power))
			if leading_zero && ((!(leading_zero_start_from_sec && !first_non0))) {
				formatString := fmt.Sprintf("%%0%d.0f%%v", int(math.Round(powerDifference)))
				// fmt.Println(formatString)
				output.WriteString(fmt.Sprintf(formatString,fl_round_time, value.Symbol+"s"))
			} else {
				output.WriteString(fmt.Sprintf("%v%v",fl_round_time, value.Symbol+"s"))
			}
			
			fl_time = fl_time - (fl_round_time * value.Base10)
			output.WriteString(" ")
			first_non0 = true
		}



		if key == break_prefix {
			break
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
	var sb strings.Builder
	// fmt.Println(input)
	// runel := []rune(input)

	for _, i := range input {
		if unicode.IsDigit(i) || i == '-' {
			sb.WriteRune(i)
		} else {
			break
		}
	}

	return strconv.ParseInt(sb.String(), 10, 64)
}


func parse_prefix_sec(input string) int64 {
	var total int64 = 0

	split := strings.SplitAfter(strings.ReplaceAll(input, " ", ""), "s")

	// fmt.Println(split)





	var rune_list []rune

	for _, i := range split {
		num, err := findAndParseNumber(i)
		if err != nil {
			// log.Println(err)
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




var show_all_values bool
var show_all_values_super bool

var leading_zero_start_from_sec bool



var round_power int64



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

	var use_all_prefixes bool

	var prefix_to_use *map[string]Prefix
	var debug bool
	var startTime time.Time

	var last_prefix string

	var last_prefix_override string

	var benchmark bool



	// Set up the command line flags
	pflag.Int64P("int_second", "i", 0, "integer second input, e.g. 1709999172")
	pflag.StringVarP(&prefix_second, "prefix_second", "p", "", "input seconds with prefixes, e.g. 1Gs 709Ms 999ks 57s")


	pflag.BoolVarP(&millisecflag, "milli", "m", false, "milliseconds")
	pflag.BoolVarP(&microsecflag, "micro", "6", false, "microseconds (6 is for 10^-6 what micro stands for)")
	pflag.BoolVarP(&nanosecflag, "nano", "n", false, "nanoseconds")

	pflag.BoolVarP(&baresecflag, "bare", "b", false, "bare integer seconds output")

	pflag.StringVarP(&date, "date", "d", "", "date input, yyyy/mm/dd HH:mm(:ss)")

	pflag.BoolVarP(&date_out, "date_out", "o", false, "date output")

	pflag.BoolVarP(&leading_zero, "leading_zeros", "l", false, "leading zeros for prefix output")
	pflag.BoolVarP(&leading_zero_start_from_sec, "leading_zeros_start_from_sec", "s", false, "disables leading zeros for first prefix")

	pflag.BoolVarP(&use_all_prefixes, "all-prefixes", "a", false, "use all prefixes instead of just common ones with a difference of 10^3")

	// pflag.BoolVarP(&show_all_values, "show_all_values", "s", false, "show_all_values even 0 ones up to last_prefix excluding 0 values before the first greater then 1")
	var hide_all_val bool
	pflag.BoolVarP(&hide_all_val, "hide_zero", "h", false, "hide zero values inside the main block. e.g. 1Gs 154Ms 0ks 54s -> 1Gs 154Ms 54s")
	pflag.BoolVarP(&show_all_values_super, "show_all_values_even_aller", "S", false, "show_all_values even 0 ones up to last_prefix including 0 values before the first greater then 1")

	pflag.BoolVar(&debug, "dbg", false, "debug")

	pflag.StringVarP(&last_prefix_override, "last_prefix_override", "f", "none", "override the last prefix to use. e.g. milli the last prefix you'll see is milli. note: none does not equal blank none means stop at no prefix")

	pflag.Int64VarP(&round_power, "round", "r", 0, "rounds down, the number is the power of 10 to round to, e.g. 1 rounds to nearest 10, 2 rounds to nearest 100")

	pflag.BoolVar(&benchmark, "ben", false, "benchmark")


	pflag.Parse()

	if debug {
		startTime = time.Now()
	}

	if hide_all_val {
		show_all_values = false
	} else {
		show_all_values = true
	}

	if use_all_prefixes {
		prefix_to_use = &AllPrefixes
	} else {
		prefix_to_use = &common_prefixes
	}

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

	last_prefix = "none"

	if millisec {
		epochTime = currentTime.UnixMilli()
		last_prefix = "milli"
		*(mul) = math.Pow(10, -3)
		// time.Unix()
	} else if microsec {
		epochTime = currentTime.UnixMicro()
		last_prefix = "micro"
		*(mul) = math.Pow(10, -6)
	} else if nanosec {
		epochTime = currentTime.UnixNano()
		last_prefix = "nano"
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

	if date != "" {
		// date, err := time.Parse(customLayout, date)
		// format := "%Y/%m/%d %H:%M"
		// date_p, err := strftime.Parse(date, format)
		// parsed_date, err := dateparse.ParseStrict(date)
		// if err != nil {
		// 	fmt.Println("err:", err)
		// }

		parsed_date := parse_date(date)
		
		// fmt.Println(parsed_date)
		// fmt.Println(parsed_date.Unix())
		if millisec {
			epochTime = parsed_date.UnixMilli()
			last_prefix = "milli"
			*(mul) = math.Pow(10, -3)
		} else if microsec {
			epochTime = parsed_date.UnixMicro()
			last_prefix = "micro"
			*(mul) = math.Pow(10, -6)
		} else if nanosec {
			epochTime = parsed_date.UnixNano()
			last_prefix = "nano"
			*(mul) = math.Pow(10, -9)
		} else {
			epochTime = parsed_date.Unix()
			mul = nil
		}
		utime = &epochTime
	} 
	if date_out {
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
		if baresec{
			fmt.Println((*utime))
		} else {
			if last_prefix_override != "none" {
				last_prefix = last_prefix_override
			}
			fmt.Println(fmt_epoch_to_prefixsec((*utime), prefix_to_use, last_prefix, mul))
		}
		
	}

	if benchmark {
		for i := 0; i < 100000; i++ {
			// Your loop body code goes here
			fmt_epoch_to_prefixsec((*utime), prefix_to_use, last_prefix, mul)
		}
	}


	if debug {
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)

		fmt.Printf("Execution time: %s\n", elapsedTime)
	}

}

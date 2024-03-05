package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

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




func fmt_epoch_to_prefixsec(utime int64, prefixes map[string]Prefix, break_prefix string) string {
	var output strings.Builder
	fl_time := float64(utime)
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


func main() {

	var utime *int64

	// Set up the command line flags
	pflag.Int64P("utime", "i", 0, "Specify the utime value")
	pflag.Parse()

	// Bind the viper configuration to the command line flags
	viper.BindPFlags(pflag.CommandLine)

	// Get the utime value from the configuration
	if viper.IsSet("utime") {
		utimeValue := viper.GetInt64("utime")
		utime = &utimeValue
	} else {
		utime = nil
	}

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the Unix epoch time in seconds
	epochTime := currentTime.Unix()

	if utime == nil {
		fmt.Println("utime is not assigned. Using default value.")
		utime = &epochTime
	}



	fmt.Println(fmt_epoch_to_prefixsec((*utime), common_prefixes, "milli"))

}

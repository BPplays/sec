package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
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




func fmt_epoch_to_prefixsec(time int64, prefixes map[string]Prefix, break_prefix string) string {
	var output strings.Builder
	fl_time := float64(time)


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
		} else {
			output.WriteString(" ")
		}


		// output.WriteString(value.FullName)

		if fl_time / value.Base10 >= 1 {
			// _ = 1
			output.WriteString(fmt.Sprintf("%v %v",fl_time / value.Base10, value.Symbol+"s"))
		}

	}

	return output.String()
}


func main() {
	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Get the Unix epoch time in seconds
	epochTime := currentTime.Unix()

	fmt.Println(fmt_epoch_to_prefixsec(epochTime, common_prefixes, "milli"))
}

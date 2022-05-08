package util

import (
	"fmt"
	"net/url"
	"strconv"
)

// ConvertTimeValuesToSeconds returns the number of seconds represented by the TimeValues
func ConvertTimeValuesToSeconds(values TimeValues) int64 {
	sum := int64(0)
	sum += values.Seconds
	sum += (values.Minutes * 60)
	sum += (values.Hours * 3600)
	sum += (values.Days * 86400)
	sum += (values.Weeks * 604800)
	sum += (values.Months * 2628000)
	return sum
}

// ConvertSecondsToTimeValues returns the number of TimeValues represented by the Seconds
func ConvertSecondsToTimeValues(seconds int64) TimeValues {
	var values TimeValues
	if seconds > 2628000 {
		values.Months = (seconds / 2628000)
		seconds -= (values.Months * 2628000)
	}
	if seconds > 604800 {
		values.Weeks = (seconds / 604800)
		seconds -= (values.Weeks * 604800)
	}
	if seconds > 86400 {
		values.Days = (seconds / 86400)
		seconds -= (values.Days * 86400)
	}
	if seconds > 3600 {
		values.Hours = (seconds / 3600)
		seconds -= (values.Hours * 3600)
	}
	if seconds > 60 {
		values.Minutes = (seconds / 60)
		seconds -= (values.Minutes * 60)
	}
	values.Seconds = seconds
	return values
}

// ConvertTimeValuesToString returns a string representation of the TimeValues
func ConvertTimeValuesToString(values TimeValues) string {
	var retString string
	if values.Months > 0 {
		retString = fmt.Sprintf("%d months", values.Months)
	}
	if values.Weeks > 0 {
		retString = fmt.Sprintf("%s %d weeks", retString, values.Weeks)
	}
	if values.Days > 0 {
		retString = fmt.Sprintf("%s %d days", retString, values.Days)
	}
	if values.Hours > 0 {
		retString = fmt.Sprintf("%s %d hours", retString, values.Hours)
	}
	if values.Minutes > 0 {
		retString = fmt.Sprintf("%s %d minutes", retString, values.Minutes)
	}
	if values.Seconds > 0 {
		retString = fmt.Sprintf("%s %d seconds", retString, values.Seconds)
	}

	return retString
}

// TimeValues represents the number of seconds, minutes, hours, days, weeks, and months
type TimeValues struct {
	Seconds int64
	Minutes int64
	Hours   int64
	Days    int64
	Weeks   int64
	Months  int64
}

// ConvertStockURLValuesToTimeValues returns the number of TimeValues represented by the StockUrlValues
func ConvertStockURLValuesToTimeValues(values url.Values) TimeValues {
	seconds, _ := strconv.ParseInt(values.Get("stockSeconds"), 10, 64)
	minutes, _ := strconv.ParseInt(values.Get("stockMinutes"), 10, 64)
	hours, _ := strconv.ParseInt(values.Get("stockHours"), 10, 64)
	days, _ := strconv.ParseInt(values.Get("stockDays"), 10, 64)
	weeks, _ := strconv.ParseInt(values.Get("stockWeeks"), 10, 64)
	months, _ := strconv.ParseInt(values.Get("stockMonths"), 10, 64)

	ret := TimeValues{
		Seconds: seconds,
		Minutes: minutes,
		Hours:   hours,
		Days:    days,
		Weeks:   weeks,
		Months:  months,
	}
	return ret
}

// ConvertPriceURLValuesToTimeValues returns the number of TimeValues represented by the PriceUrlValues
func ConvertPriceURLValuesToTimeValues(values url.Values) TimeValues {
	seconds, _ := strconv.ParseInt(values.Get("wlSeconds"), 10, 64)
	minutes, _ := strconv.ParseInt(values.Get("wlMinutes"), 10, 64)
	hours, _ := strconv.ParseInt(values.Get("wlHours"), 10, 64)
	days, _ := strconv.ParseInt(values.Get("wlDays"), 10, 64)
	weeks, _ := strconv.ParseInt(values.Get("wlWeeks"), 10, 64)
	months, _ := strconv.ParseInt(values.Get("wlMonths"), 10, 64)

	ret := TimeValues{
		Seconds: seconds,
		Minutes: minutes,
		Hours:   hours,
		Days:    days,
		Weeks:   weeks,
		Months:  months,
	}
	return ret
}

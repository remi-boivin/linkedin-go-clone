package util

import (
	"net/url"
	"strconv"
)

// ConvertStringFieldsToInt returns a map of string keys to int values
func ConvertStringFieldsToInt(values url.Values, keys []string) map[string]int {
	ret := map[string]int{}
	for _, key := range keys {
		value := values.Get(key)

		intValue, _ := strconv.Atoi(value)
		ret[key] = intValue
	}
	return ret
}

// ConvertStringFieldsToFloat returns a map of string keys to float values
func ConvertStringFieldsToFloat(values url.Values, keys []string) map[string]float32 {
	ret := map[string]float32{}
	for _, key := range keys {
		value := values.Get(key)

		intValue, _ := strconv.ParseFloat(value, 32)
		ret[key] = float32(intValue)
	}
	return ret
}

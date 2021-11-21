package util

import (
	"strconv"
)

/* string
   int
   float64
   []byte
*/

// StringToInt string => int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// IntToString int => string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// IntToInt64 int => int64
func IntToInt64(i int) int64 {
	return int64(i)
}

// StringToFloat string => int, 注意这个bitSize要么是32，要么是64，如果瞎写的话默认给你算64
func StringToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return f
}

// FloatToString float64 => string
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

// StringToBytes string => byte
func StringToBytes(s string) []byte {
	return []byte(s)
}

// BytesToString string => byte
func BytesToString(b []byte) string {
	return string(b)
}

// IntToFloat64 Int => float64
func IntToFloat64(i int) float64 {
	return float64(i)
}

// Float64ToInt IntToFloat64 Int => float64
func Float64ToInt(f float64) int {
	return int(f)
}

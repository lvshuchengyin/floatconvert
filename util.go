package gobignumber

import (
	"strconv"
	"strings"
)

// RightPad2Len add padstr to extend string length
func RightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// LeftPad2Len add padstr to extend string length
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func IsValidFloat(s string) (v bool) {
	if strings.Contains(s, "e") || strings.Contains(s, "E") || strings.HasPrefix(s, ".") {
		v = false
		return
	}
	_, err := strconv.ParseFloat(s, 32)
	v = err == nil
	return
}

// IsValidInt64 not support scientific notation
func IsValidInt64(s string) (v bool) {
	if strings.Contains(s, "e") || strings.Contains(s, "E") {
		v = false
		return
	}

	_, err := strconv.ParseInt(s, 10, 64)
	v = err == nil
	return
}

func Pow10(n int) (v int64) {
	if n <= 0 {
		v = 1
		return
	}

	v = 1
	for i := 0; i < n; i++ {
		v *= 10
	}

	return
}

// StringToInt64 convert string to int64
func StringToInt64(s string, defaultVal int64) (i int64) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = defaultVal
		return
	}
	return
}

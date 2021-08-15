package gofloatconvert

import (
	"strings"
)

// FloatStringToInt64 convert float string to 10 * decimalPlaces int64, in case of precision problem
func FloatStringToInt64(s string, decimalPlaces int) (i int64, err error) {
	s = strings.TrimSpace(s)

	if decimalPlaces < 0 {
		err = ErrInvalidDecimalPlaces
		return
	}
	multiples := Pow10(decimalPlaces)

	if !IsValidFloat(s) {
		err = ErrInvalidFloatString
		return
	}

	if !strings.Contains(s, ".") {
		i = StringToInt64(s, 0) * multiples
		return
	}

	isPositive := true
	if strings.HasPrefix(s, "-") {
		isPositive = false
	}

	parts := strings.Split(s, ".")
	integerPart := parts[0]
	decimalPart := parts[1]
	decimalPart = RightPad2Len(decimalPart, "0", decimalPlaces)

	i = StringToInt64(integerPart, 0) * multiples
	if isPositive {
		i += StringToInt64(decimalPart, 0)
	} else {
		i -= StringToInt64(decimalPart, 0)
	}
	return
}

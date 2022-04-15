package gobignumber

import (
	"fmt"
	"strconv"
	"strings"
)

// Int64ToFloatString convert int64 10w to float string
func Int64ToFloatString(i int64, decimalPlaces int) (s string, err error) {
	isPositive := true
	if i < 0 {
		isPositive = false
		i = -i
	}

	if decimalPlaces < 0 {
		err = ErrInvalidDecimalPlaces
		return
	}
	multiples := Pow10(decimalPlaces)

	integerPart := i / multiples
	decimalPart := i % multiples

	integerPartStr := strconv.FormatInt(integerPart, 10)
	if decimalPart == 0 {
		s = integerPartStr
		if !isPositive {
			s = "-" + s
		}
		return
	}

	decimalPartStr := LeftPad2Len(strconv.FormatInt(decimalPart, 10), "0", 5)
	s = fmt.Sprintf("%s.%s", integerPartStr, decimalPartStr)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	if !isPositive {
		s = "-" + s
	}
	return
}

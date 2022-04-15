package gobignumber

import (
	"errors"
)

var (
	ErrInvalidFloatString   = errors.New("invalid float string")
	ErrInvalidDecimalPlaces = errors.New("invalid decimal places")
)

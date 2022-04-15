package gobignumber

import (
	"errors"
	"math/big"
	"strings"
)

type BigNumber struct {
	bigint        *big.Int
	decimalPlaces int
}

func (b *BigNumber) String() (res string) {
	if b.decimalPlaces <= 0 {
		return b.bigint.String()
	}

	innerBStr := b.bigint.String()
	isPositive := !strings.HasPrefix(innerBStr, "-")
	innerBStr = strings.TrimLeft(innerBStr, "-")

	if len(innerBStr) < b.decimalPlaces {
		res = "0." + LeftPad2Len(innerBStr, "0", b.decimalPlaces)
	} else if len(innerBStr) == b.decimalPlaces {
		res = "0." + innerBStr
	} else {
		res = innerBStr[0:len(innerBStr)-b.decimalPlaces] + "." + innerBStr[len(innerBStr)-b.decimalPlaces:]
	}

	if !isPositive {
		res = "-" + res
	}
	return
}

func (b *BigNumber) Add(x *BigNumber) (res *BigNumber) {
	maxDecimalPlaces := b.decimalPlaces
	if x.decimalPlaces > maxDecimalPlaces {
		maxDecimalPlaces = x.decimalPlaces
	}

	bInnerB := BigIntMul10(b.bigint, maxDecimalPlaces-b.decimalPlaces)
	bInnerX := BigIntMul10(x.bigint, maxDecimalPlaces-x.decimalPlaces)

	res = &BigNumber{
		bigint:        big.NewInt(0).Add(bInnerB, bInnerX),
		decimalPlaces: maxDecimalPlaces,
	}
	return
}

func (b *BigNumber) Sub(x *BigNumber) (res *BigNumber) {
	maxDecimalPlaces := b.decimalPlaces
	if x.decimalPlaces > maxDecimalPlaces {
		maxDecimalPlaces = x.decimalPlaces
	}

	bInnerB := BigIntMul10(b.bigint, maxDecimalPlaces-b.decimalPlaces)
	bInnerX := BigIntMul10(x.bigint, maxDecimalPlaces-x.decimalPlaces)

	res = &BigNumber{
		bigint:        big.NewInt(0).Sub(bInnerB, bInnerX),
		decimalPlaces: maxDecimalPlaces,
	}
	return
}

func (b *BigNumber) Mul(x *BigNumber) (res *BigNumber) {
	bInnerB := b.bigint
	bInnerX := x.bigint

	res = &BigNumber{
		bigint:        big.NewInt(0).Mul(bInnerB, bInnerX),
		decimalPlaces: b.decimalPlaces + x.decimalPlaces,
	}
	return
}

func (b *BigNumber) Quo(x *BigNumber, decimalPlaces int) (res *BigNumber) {
	maxDecimalPlaces := b.decimalPlaces
	if x.decimalPlaces > maxDecimalPlaces {
		maxDecimalPlaces = x.decimalPlaces
	}

	bInnerB := BigIntMul10(b.bigint, maxDecimalPlaces-b.decimalPlaces+decimalPlaces)
	bInnerX := BigIntMul10(x.bigint, maxDecimalPlaces-x.decimalPlaces)

	res = &BigNumber{
		bigint:        big.NewInt(0).Quo(bInnerB, bInnerX),
		decimalPlaces: decimalPlaces,
	}
	return
}

func NewBigNumber(n int64) *BigNumber {
	return &BigNumber{
		bigint:        big.NewInt(n),
		decimalPlaces: 0,
	}
}

func NewBigNumberFromString(s string) (b *BigNumber, err error) {
	s = strings.TrimSpace(s)

	if !IsValidFloat(s) {
		err = errors.New("not a number string")
		return
	}

	if !strings.Contains(s, ".") {
		innerB, ok := new(big.Int).SetString(s, 10)
		if !ok {
			err = errors.New("not a valid number string")
			return
		}
		b = &BigNumber{
			bigint:        innerB,
			decimalPlaces: 0,
		}
		return
	}

	isPositive := true
	if strings.HasPrefix(s, "-") {
		isPositive = false
	}

	parts := strings.Split(s, ".")
	integerPart := parts[0]
	decimalPart := parts[1]
	decimalPart = strings.TrimRight(decimalPart, "0")

	integerPartInnerB, ok := new(big.Int).SetString(integerPart, 10)
	if !ok {
		err = errors.New("integer part not a valid number string")
		return
	}

	decimalPartInnerB, ok := new(big.Int).SetString(decimalPart, 10)
	if !ok {
		err = errors.New("decimal part not a valid number string")
		return
	}

	decimalPlaces := len(decimalPart)

	for i := 0; i < decimalPlaces; i++ {
		integerPartInnerB = big.NewInt(0).Mul(integerPartInnerB, big.NewInt(10))
	}

	var innerB *big.Int
	if isPositive {
		innerB = big.NewInt(0).Add(integerPartInnerB, decimalPartInnerB)
	} else {
		innerB = big.NewInt(0).Sub(integerPartInnerB, decimalPartInnerB)
	}

	b = &BigNumber{
		bigint:        innerB,
		decimalPlaces: decimalPlaces,
	}
	return
}

func BigIntMul10(bi *big.Int, n int) (res *big.Int) {
	res = big.NewInt(0).Set(bi)
	for i := 0; i < n; i++ {
		res = big.NewInt(0).Mul(res, big.NewInt(10))
	}
	return
}

package main

import (
	"fmt"

	"github.com/lvshuchengyin/gofloatconvert"
)

const (
	DecimalPlaces = 5
)

func main() {
	floatString := "123.456789"
	int64WithPricision := int64(0)
	int64WithPricision, err := gofloatconvert.FloatStringToInt64(floatString, DecimalPlaces)
	if err != nil {
		fmt.Println("FloatStringToInt64 err", err.Error())
		return
	}
	fmt.Printf("FloatStringToInt64, floatString:%v, int64WithPricision:%v\n", floatString, int64WithPricision)

	floatString, err = gofloatconvert.Int64ToFloatString(int64WithPricision, DecimalPlaces)
	if err != nil {
		fmt.Println("Int64ToFloatString err", err.Error())
		return
	}
	fmt.Printf("Int64ToFloatString, int64WithPricision:%v, floatString:%v\n", int64WithPricision, floatString)
}

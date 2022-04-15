package main

import (
	"fmt"

	"github.com/lvshuchengyin/gobignumber"
)

const (
	DecimalPlaces = 5
)

func main() {
	floatString := "123.456789"
	int64WithPricision := int64(0)
	int64WithPricision, err := gobignumber.FloatStringToInt64(floatString, DecimalPlaces)
	if err != nil {
		fmt.Println("FloatStringToInt64 err", err.Error())
		return
	}
	fmt.Printf("FloatStringToInt64, floatString:%v, int64WithPricision:%v\n", floatString, int64WithPricision)

	floatString, err = gobignumber.Int64ToFloatString(int64WithPricision, DecimalPlaces)
	if err != nil {
		fmt.Println("Int64ToFloatString err", err.Error())
		return
	}
	fmt.Printf("Int64ToFloatString, int64WithPricision:%v, floatString:%v\n", int64WithPricision, floatString)

	{
		sbn, _ := gobignumber.NewBigNumberFromString("-0.00034567")
		sbn2, _ := gobignumber.NewBigNumberFromString("-12.345678")
		sbn3 := sbn.Add(sbn2)
		fmt.Println("Add", sbn, sbn2, sbn3)
	}

	{
		sbn, _ := gobignumber.NewBigNumberFromString("-0.00034567")
		sbn2, _ := gobignumber.NewBigNumberFromString("-12.345678")
		sbn3 := sbn.Sub(sbn2)
		fmt.Println("Sub", sbn, sbn2, sbn3)
	}

	{
		sbn, _ := gobignumber.NewBigNumberFromString("-0.00034567")
		sbn2, _ := gobignumber.NewBigNumberFromString("-12.345678")
		sbn3 := sbn.Mul(sbn2)
		fmt.Println("Mul", sbn, sbn2, sbn3)
	}

	{
		sbn, _ := gobignumber.NewBigNumberFromString("-0.00034567")
		sbn2, _ := gobignumber.NewBigNumberFromString("-12.345678")
		sbn3 := sbn.Quo(sbn2, 10)
		fmt.Println("Quo", sbn, sbn2, sbn3)
	}
}

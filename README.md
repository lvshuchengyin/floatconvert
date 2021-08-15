# gofloatconvert
A float convert library for Golang.

# Usage scenarios
Float is a imprecise number. So when convert float to string, will have precision loss.

In order to save float value which input by user, we may design the api to let front end to pass float string, so back end can get the value without pricision, and then transform float string to int64, finally save to DB(MySQL).

When front end get the value, back end should convert the int64 to float string.
So, we could deal with float number without precision loss.

# Overview
- float string convert to int64
- int64 with pricision convert to float string

# Example
``` go
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
```

**output:**
```
FloatStringToInt64, floatString:123.456789, int64WithPricision:12345678
Int64ToFloatString, int64WithPricision:12345678, floatString:123.45678
```

# TODO
- deal with overflow 

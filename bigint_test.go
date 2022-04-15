package gobignumber

import (
	"testing"
)

func Test_IntPointer(t *testing.T) {
	{
		sbn, _ := NewBigNumberFromString("-0.00034567")
		sbn2, _ := NewBigNumberFromString("-12.345678")
		sbn3 := sbn.Add(sbn2)
		if sbn3.String() != "-12.34602367" {
			t.Error("Add", sbn, sbn2, sbn3)
		}
	}

	{
		sbn, _ := NewBigNumberFromString("-0.00034567")
		sbn2, _ := NewBigNumberFromString("-12.345678")
		sbn3 := sbn.Sub(sbn2)
		if sbn3.String() != "12.34533233" {
			t.Error("Sub", sbn, sbn2, sbn3)
		}
	}

	{
		sbn, _ := NewBigNumberFromString("-0.00034567")
		sbn2, _ := NewBigNumberFromString("-12.345678")
		sbn3 := sbn.Mul(sbn2)
		if sbn3.String() != "0.00426753051426" {
			t.Error("Mul", sbn, sbn2, sbn3)
		}
	}

	{
		sbn, _ := NewBigNumberFromString("-0.00034567")
		sbn2, _ := NewBigNumberFromString("-12.345678")
		sbn3 := sbn.Quo(sbn2, 10)
		if sbn3.String() != "0.0000279992" {
			t.Error("Quo", sbn, sbn2, sbn3)
		}
	}
}

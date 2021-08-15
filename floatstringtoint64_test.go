package gofloatconvert

import (
	"strconv"
	"testing"
)

func Test_FloatStringToInt64(t *testing.T) {
	var datas = [][]string{
		[]string{"1", "100000"},
		[]string{"1.01", "101000"},
		[]string{"1.", "100000"},
		[]string{"1..", "0"},
		[]string{"1.a", "0"},
		[]string{".1", "10000"},
		[]string{"100", "10000000"},
		[]string{"1.23456789123456", "123456"},
		[]string{"0.0001", "10"},
		[]string{"0.0000000001", "0"},
		[]string{"123", "12300000"},
		[]string{"123.456789", "12345678"},
		[]string{"a", "0"},
		[]string{"", "0"},
		[]string{" 123", "12300000"},
		[]string{"-123", "-12300000"},
		[]string{"-1.23456789123456", "-123456"},
		[]string{"-0.0001", "-10"},
	}
	for _, item := range datas {
		in := item[0]
		expect, _ := strconv.ParseInt(item[1], 10, 64)
		r, _ := FloatStringToInt64(in, 5)
		if r != expect {
			t.Errorf("error, in: %v, out:%v, expect:%v", in, r, expect)
		}
	}
}

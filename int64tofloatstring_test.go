package gofloatconvert

import (
	"strconv"
	"testing"
)

func Test_Int64ToFloatString(t *testing.T) {
	var datas = [][]string{
		[]string{"123", "0.00123"},
		[]string{"100000", "1"},
		[]string{"0", "0"},
		[]string{"123456789", "1234.56789"},
		[]string{"100", "0.001"},
		[]string{"1", "0.00001"},
		[]string{"1000000", "10"},
		[]string{"-123", "-0.00123"},
		[]string{"-100000", "-1"},
		[]string{"-0", "0"},
		[]string{"-123456789", "-1234.56789"},
		[]string{"-100", "-0.001"},
		[]string{"-1", "-0.00001"},
		[]string{"-1000000", "-10"},
	}
	for _, item := range datas {
		in, _ := strconv.ParseInt(item[0], 10, 64)
		expect := item[1]
		r, err := Int64ToFloatString(in, 5)
		if err != nil {
			t.Errorf("error, in: %v, err:%v", in, err)
		}
		if r != expect {
			t.Errorf("error, in: %v, out:%v, expect:%v", in, r, expect)
		}
	}
}

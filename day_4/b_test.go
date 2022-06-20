package day_4


import (
	"github.com/shopspring/decimal"
	"testing"
)

func Ff(x float64) {
	decimal.NewFromFloat(x).Round(2).Float64()
	//fmt.Println(v1)
}

func BenchmarkFf(b *testing.B) {
	for n := 0; n < b.N; n++{
		Ff(9.7531)
	}
}

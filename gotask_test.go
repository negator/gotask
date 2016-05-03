package gotask

import (
	"fmt"
	"math"
	"testing"
)

func TestGoTaskGroup(b *testing.T) {

	a := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 9, 9, 10, 11, 12}

	batchSize := 3
	tg := GoTaskGroup{}
	c := make(chan []int, 1000)
	for i := 0; i < len(a); i += batchSize {
		j := i
		tg.AppendF(func() {
			s := math.Min(float64(len(a)), float64(j+batchSize))
			b := a[j:int(s)]
			c <- b
			fmt.Println(b)
		})
	}
	tg.AwaitAll()
}

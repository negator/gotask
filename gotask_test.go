package gotask

import (
	"fmt"
	"math"
	"testing"
)

func TestGoTaskGroup(b *testing.T) {

	a := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 9, 9, 10, 11, 12}

	batchSize := 3
	var tg GoTaskGroup
	c := make(chan []int, 1000)
	defer close(c)
	for i := 0; i < len(a); i += batchSize {
		j := i
		tg = tg.AppendF(func() {
			s := math.Min(float64(len(a)), float64(j+batchSize))
			b := a[j:int(s)]
			c <- b
			fmt.Println(b)
		})
	}
	tg.AwaitAll()
}

func TestGoTask(b *testing.T) {
	c := make(chan int, 1)

	producer := goTask(func() {
		/*Potentially long running IO like a streaming database call.*/
		for i := 0; i < 100; i++ {
			var j = i
			c <- j
		}
		close(c)
	})

	consumer := goTask(func() {
		for i := range c {
			fmt.Printf("Consumed %v\n", i)
		}
	})
	var tg GoTaskGroup
	tg = tg.Append(producer, consumer)
	tg.AwaitAll()
}

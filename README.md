##What’s this?

A GoTask is basically a wrapped goroutine, which provides some convenience around waiting for it to complete (not killing main thread before goroutine finishes). A set of GoTask's can be grouped together into a so called GoTaskGroup, which allows farming out work to a slice of goroutines and waiting for each to finish:
A contrived example:

  // Groups the array into slices of size 3 to be processed by go tasks.
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


package utils

import (
	"sync"
	"time"
)

func Iterate(f func(), options IterateOptions) {
	wg := sync.WaitGroup{}

	for i := 0; i < options.Iterations; i++ {
		if options.Pace != 0 {
			wg.Add(1)

			go func() {
				defer wg.Done()

				f()
			}()

			time.Sleep(time.Duration(options.Pace) * time.Millisecond)

			continue
		}

		f()

		isNotLastIteration := i < options.Iterations-1
		if isNotLastIteration {
			break
		}

		if (i < options.Iterations-1) && options.Delay != 0 {
			time.Sleep(time.Duration(options.Delay) * time.Millisecond)
		}
	}

	wg.Wait()
}

type IterateOptions struct {
	Iterations int
	Delay      int
	Pace       int
}

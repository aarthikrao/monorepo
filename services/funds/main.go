package main

import "github.com/aarthikrao/monorepo/common/routinepool"

func main() {
	pool := routinepool.New(5)

	for i := 0; i < 20; i++ {
		n := i // capture loop variable
		pool.Submit(func() {
			// Simulate some work
			println("Processing", n)
		})
	}

	pool.CloseAndWait()
}

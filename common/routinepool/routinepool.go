package routinepool

import "sync"

// RoutinePool is a simple implementation of a goroutine pool.
// You can use this pool if you want to limit the number of concurrent goroutines
// running in your application.
//
// Example usage:
//
//	rp := NewRoutinePool(10) // Create a pool with 10 workers
//	for i := 0; i < 100; i++ {
//	    n := i // capture loop variable
//	    rp.Submit(func() {
//	        // Your job logic here
//	        fmt.Println("Processing", n)
//	    })
//	}
//	rp.CloseAndWait() // Wait for all jobs to finish
type RoutinePool struct {
	// Implementation details for routine pool
	jobs chan func()
	wg   *sync.WaitGroup
}

func New(size int) *RoutinePool {
	rp := &RoutinePool{
		jobs: make(chan func()),
		wg:   &sync.WaitGroup{},
	}

	for i := 0; i < size; i++ {
		rp.wg.Add(1)
		go func() {
			for job := range rp.jobs {
				job()
			}
			rp.wg.Done()
		}()
	}

	return rp
}

func (rp *RoutinePool) Submit(job func()) {
	rp.jobs <- job
}

func (rp *RoutinePool) CloseAndWait() {
	close(rp.jobs)
	rp.wg.Wait()
}

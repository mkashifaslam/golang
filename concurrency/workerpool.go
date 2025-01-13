package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID     int
	Processed string
	Error     error
}

func WorkerPool(jobs []Job, numWorkers int) []Result {
	var wg sync.WaitGroup
	jobsChan := make(chan Job, len(jobs))       // Buffered channel for jobs
	resultsChan := make(chan Result, len(jobs)) // Buffered channel for results

	// start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		fmt.Println("starting worker", i+1)
		go worker(&wg, jobsChan, resultsChan)
	}

	fmt.Println("All workers started")

	// Send jobs
	go func() {
		for _, job := range jobs {
			jobsChan <- job
		}
		fmt.Println("All jobs sent")
		close(jobsChan)
	}()

	// Wait and collect results
	go func() {
		wg.Wait()
		fmt.Println("All workers finished")
		close(resultsChan)
	}()

	var results []Result
	for result := range resultsChan {
		results = append(results, result)
	}

	fmt.Println("All workers results compiled")

	return results

}

// worker jobs is receive-only channel and results is send-only channel
func worker(wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("job run:", job.ID)
		result := Result{
			JobID:     job.ID,
			Processed: "processed-" + job.Data,
		}
		results <- result
	}
}

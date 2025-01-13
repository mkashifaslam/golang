package main

import "fmt"

func main() {
	jobs := []Job{
		{ID: 1, Data: "id"},
		{ID: 2, Data: "title"},
		{ID: 3, Data: "status"},
		{ID: 4, Data: "description"},
		{ID: 5, Data: "time"},
	}

	result := WorkerPool(jobs, 5)

	for _, result := range result {
		fmt.Println("job", result.JobID)
		fmt.Println("result", result.Processed)
	}

}

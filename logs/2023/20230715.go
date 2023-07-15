package main

import (
	"fmt"
)

type Job struct {
	Name string
}

func countJobs(jobs []Job) {
	fmt.Println(len(jobs))
	fmt.Println(jobs)
}

func main() {
	jobs := []Job{
		{Name: "Job 1"},
		{Name: "Job 2"},
		{Name: "Job 3"},
	}
	countJobs(jobs)
	noJobs := []Job{}
	countJobs(noJobs)
	countJobs(nil)
}

// Output:
// 3
// [{Job 1} {Job 2} {Job 3}]
// 0
// []
// 0
// []
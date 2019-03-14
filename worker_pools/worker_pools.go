// this program shows the core functionalities of worker groups

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	randomNumber int
}

type Result struct {
	job Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)

	return sum
}

func worker (wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{
			job,
			digits(job.randomNumber),
		}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(numberOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(numberOfJobs int) {
	for i := 0; i < numberOfJobs; i++ {
		randomValue := rand.Intn(999)
		job := Job{
			i,
			randomValue,
		}
		jobs <- job
	}
	close(jobs)
}

func result (done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random number %d, sum of digits %d\n", result.job.id, result.job.randomNumber, result.sumOfDigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()

	numberOfJobs := 100
	go allocate(numberOfJobs)

	done := make(chan bool)
	go result(done)

	numberOfWorkers := 10
	go createWorkerPool(numberOfWorkers)
	
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Time taken ", diff.Seconds(), "seconds")
}
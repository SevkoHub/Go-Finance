package worker

import (
	"fmt"
	"sync"
)

// WorkerPool struct to hold the workers and the job queue
type WorkerPool struct {
	workers  []*Worker
	jobQueue chan func() // Channel to receive jobs
	wg       sync.WaitGroup
}

// Worker struct represents a single worker
type Worker struct {
	id      int
	pool    *WorkerPool
}

// NewWorkerPool creates a new WorkerPool with a specified number of workers
func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{
		jobQueue: make(chan func()), // Initialize the job queue
	}

	// Create workers and start them
	for i := 0; i < numWorkers; i++ {
		worker := &Worker{id: i + 1, pool: pool}
		pool.workers = append(pool.workers, worker)
		go worker.start() // Start the worker in a goroutine
	}

	return pool
}

// Start method for the worker to listen for jobs
func (w *Worker) start() {
	for job := range w.pool.jobQueue { // Listen for jobs in the job queue
		job() // Execute the job
		w.pool.wg.Done() // Mark the job as done
	}
}

// Submit method to add a job to the worker pool
func (p *WorkerPool) Submit(job func()) {
	p.wg.Add(1) // Increment the wait group counter
	p.jobQueue <- job // Send the job to the job queue
}

// Shutdown method to gracefully stop the worker pool
func (p *WorkerPool) Shutdown() {
	close(p.jobQueue) // Close the job queue to stop workers
	p.wg.Wait() // Wait for all jobs to finish
	fmt.Println("All workers have completed their tasks.")
}
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

func worker(ctx context.Context, id int, jobs <-chan Job, limiter <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		select {
		case <-limiter:
			fmt.Printf("worker %d mengerjakan job %d: %s\n", id, job.ID, job.Payload)
		case <-ctx.Done():
			fmt.Printf("worker %d berhenti: %v\n", id, ctx.Err())
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	jobs := make(chan Job, 12)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, ticker.C, &wg)
	}

	for i := 1; i <= 12; i++ {
		jobs <- Job{ID: i, Payload: fmt.Sprintf("kerjakan job ke %d", i)}
	}
	close(jobs)

	wg.Wait()
	fmt.Println("tugas selesai!")
}

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

func worker(
	ctx context.Context,
	id int,
	jobs <-chan Job,
	limiter <-chan time.Time,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		// Tahap 1: Menunggu job atau cancellation
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d berhenti: %v\n", id, ctx.Err())
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d selesai: antrean kosong\n", id)
				return
			}

			// Tahap 2: Menunggu izin atau cancellation
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d membatalkan Job %d\n", id, job.ID)
				return

			case <-limiter:
				// Periksa lagi sebelum memulai pekerjaan.
				if ctx.Err() != nil {
					return
				}

				fmt.Printf(
					"Worker %d mengerjakan Job %d: %s\n",
					id,
					job.ID,
					job.Payload,
				)
			}
		}
	}
}

func main() {

	jobs := make(chan Job, 12)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	// 3 worker
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, ticker.C, &wg)
	}

	// 12 buffer

	for i := 1; i <= 12; i++ {
		jobs <- Job{
			ID:      i,
			Payload: fmt.Sprintf("mengerjakan job ke %d", i),
		}
	}

	close(jobs)

	wg.Wait()

	fmt.Println("tugas selesai!!!")

}

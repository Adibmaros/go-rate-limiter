package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
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
	succeeded *atomic.Int64,
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

			// Tahap 2: Menunggu izin rate limiter atau cancellation
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d membatalkan Job %d\n", id, job.ID)
				return

			case <-limiter:
				// Periksa kembali apakah context sudah timeout.
				if ctx.Err() != nil {
					return
				}

				fmt.Printf(
					"Worker %d mengerjakan Job %d: %s\n",
					id,
					job.ID,
					job.Payload,
				)

				// Simulasi pekerjaan.
				// Pekerjaan dapat dibatalkan jika context timeout.
				select {
				case <-time.After(200 * time.Millisecond):
					// Job benar-benar selesai.
					succeeded.Add(1)

				case <-ctx.Done():
					fmt.Printf("Worker %d gagal menyelesaikan Job %d\n", id, job.ID)
					return
				}
			}
		}
	}
}

func main() {
	const totalJobs = 12

	jobs := make(chan Job, totalJobs)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		4*time.Second,
	)
	defer cancel()

	var wg sync.WaitGroup
	var succeeded atomic.Int64

	// Menjalankan 3 worker.
	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go worker(
			ctx,
			i,
			jobs,
			ticker.C,
			&wg,
			&succeeded,
		)
	}

	// Producer mengirimkan 12 job.
	for i := 1; i <= totalJobs; i++ {
		jobs <- Job{
			ID:      i,
			Payload: fmt.Sprintf("mengerjakan job ke %d", i),
		}
	}

	// Producer selesai mengirim job.
	close(jobs)

	// Tunggu semua worker selesai.
	wg.Wait()

	// Hitung hasil pekerjaan.
	completed := succeeded.Load()
	unfinished := int64(totalJobs) - completed

	fmt.Println("\n========== HASIL PEMROSESAN ==========")
	fmt.Println("Total job       :", totalJobs)
	fmt.Println("Berhasil        :", completed)
	fmt.Println("Tidak selesai   :", unfinished)
	fmt.Println("======================================")
}

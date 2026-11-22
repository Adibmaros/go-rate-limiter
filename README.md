# Go Rate Limiter & Worker Pool with Multi-Stage Cancellation

Implementasi **Worker Pool** konkuren yang dipadukan dengan **Rate Limiting** dan **Multi-Stage Graceful Cancellation** menggunakan fitur bawaan Go standard library (`goroutine`, `channel`, `time.Ticker`, `sync/atomic`, dan `context.Context`).

---

## 🎯 Gambaran Proyek

Proyek ini mendemonstrasikan pola penanganan beban kerja (*job queue*) di mana request/job diproses oleh kumpulan worker dengan batasan kecepatan eksekusi (*rate limit*), sekaligus mampu menangani pembatalan bertahap (*timeout/cancellation*) secara aman tanpa *goroutine leak* atau *data race*.

### Parameter Simulasi Default
- **Total Pekerjaan**: 12 jobs
- **Jumlah Worker**: 3 goroutine aktif
- **Rate Limit**: 1 izin setiap 500 ms (`time.NewTicker`)
- **Durasi Simulasi per Job**: 200 ms
- **Batas Waktu Global (Timeout)**: 4 detik (`context.WithTimeout`)

---

## 🚀 Fitur & Konsep Utama

1. **Worker Pool Pattern**:
   Menggunakan channel sebagai antrean job yang dikonsumsi secara paralel oleh 3 worker goroutines.
2. **Channel-based Rate Limiting**:
   Membatasi frekuensi pemrosesan job dengan membaca *tick* dari channel `time.Ticker.C`.
3. **Multi-Stage Graceful Cancellation**:
   Worker memvalidasi sinyal `ctx.Done()` di setiap tahapan kritis:
   - **Tahap 1**: Saat menunggu job dari antrean channel atau antrean sudah di-close (`ok == false`).
   - **Tahap 2**: Saat menunggu jatah/izin dari *rate limiter*.
   - **Tahap 3**: Selama proses eksekusi pekerjaan (jika timeout terjadi di tengah proses, job ditandai gagal).
4. **Thread-Safe Metrics (`sync/atomic`)**:
   Menghitung jumlah pekerjaan yang berhasil diselesaikan menggunakan `atomic.Int64` tanpa perlunya mutex manual.
5. **Ringkasan Hasil Eksekusi**:
   Menampilkan rekapitulasi status akhir (*Total job*, *Berhasil*, dan *Tidak selesai*).

---

## 📁 Struktur Proyek

```text
go-rate-limiter/
├── .gitignore       # Konfigurasi file yang diabaikan oleh Git
├── go.mod           # Konfigurasi modul Go
├── main.go          # Implementasi lengkap worker pool, rate limiter, & metrics
└── README.md        # Dokumentasi proyek
```

---

## 🛠️ Prasyarat

- [Go](https://go.dev/dl/) versi 1.21 ke atas (disarankan Go 1.25+ sesuai `go.mod`).

---

## ⚡ Cara Menjalankan

1. **Jalankan program langsung:**
   ```bash
   go run main.go
   ```

2. **Kompilasi menjadi biner:**
   ```bash
   go build -o rate-limiter.exe main.go
   ./rate-limiter.exe
   ```

---

## 📋 Contoh Output

Karena total waktu yang dibutuhkan untuk 12 job dengan jeda 500ms melebihi batas timeout 4 detik, program akan memproses sekitar 7–8 job sebelum timeout membatalkan sisa pekerjaan:

```text
Worker 1 mengerjakan Job 1: mengerjakan job ke 1
Worker 3 mengerjakan Job 2: mengerjakan job ke 2
Worker 2 mengerjakan Job 3: mengerjakan job ke 3
Worker 1 mengerjakan Job 4: mengerjakan job ke 4
Worker 3 mengerjakan Job 5: mengerjakan job ke 5
Worker 2 mengerjakan Job 6: mengerjakan job ke 6
Worker 1 mengerjakan Job 7: mengerjakan job ke 7
Worker 3 membatalkan Job 8
Worker 1 membatalkan Job 10
Worker 2 membatalkan Job 9

========== HASIL PEMROSESAN ==========
Total job       : 12
Berhasil        : 7
Tidak selesai   : 5
======================================
```

---

## 📌 Rencana Pengembangan (Roadmap)

- [x] Worker pool dengan rate limiting (`time.Ticker`).
- [x] Multi-stage timeout & cancellation (`context.Context`).
- [x] Perhitungan metrik thread-safe (`sync/atomic`).
- [ ] Implementasi algoritma Token Bucket / Leaky Bucket.
- [ ] Dukungan konfigurasi dinamis (jumlah worker, rate limit, timeout via parameter/CLI).
- [ ] Integrasi middleware HTTP (Fiber / Gin / Standard `net/http`).
- [ ] Distributed rate limiting menggunakan Redis.
- [ ] Unit testing dan benchmark concurrency.

---

## 📄 Lisensi

Proyek ini dibuat untuk tujuan pembelajaran dan eksplorasi pola konkurensi di Go. Bebas digunakan dan dimodifikasi.

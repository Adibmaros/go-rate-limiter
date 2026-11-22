# Go Rate Limiter & Worker Pool

Proyek sederhana berbasis Go untuk mendemonstrasikan implementasi **Worker Pool** yang dipadukan dengan **Rate Limiting** menggunakan fitur bawaan Go (`goroutine`, `channel`, `time.Ticker`, dan `context.Context`).

> **Catatan**: Repositori ini sedang dalam tahap pengembangan aktif dan kode akan terus diperbarui.

---

## 🚀 Fitur & Konsep Utama

- **Worker Pool**: Memproses banyak *job* secara konkuren menggunakan sejumlah worker goroutine yang terkontrol.
- **Rate Limiting**: Membatasi laju eksekusi task menggunakan `time.Ticker` (misal: 1 job per 500ms).
- **Graceful Cancellation & Timeout**: Menggunakan `context.WithTimeout` untuk menghentikan worker secara aman jika melebihi batas waktu eksekusi.
- **Sinkronisasi Goroutine**: Menggunakan `sync.WaitGroup` untuk memastikan semua worker selesai sebelum program keluar.

---

## 📁 Struktur Proyek

```text
go-rate-limiter/
├── .gitignore       # Konfigurasi file yang diabaikan oleh Git
├── go.mod           # Definisi modul Go
├── main.go          # Implementasi worker pool & rate limiter
└── README.md        # Dokumentasi proyek
```

---

## 🛠️ Prasyarat

- [Go](https://go.dev/dl/) versi 1.21 atau yang lebih baru (proyek ini menggunakan Go 1.25).

---

## ⚡ Cara Menjalankan

1. **Clone repository (jika dari remote):**
   ```bash
   git clone <url-repository>
   cd go-rate-limiter
   ```

2. **Jalankan langsung:**
   ```bash
   go run main.go
   ```

3. **Build binary executable:**
   ```bash
   go build -o rate-limiter.exe main.go
   ./rate-limiter.exe
   ```

---

## 📋 Contoh Output

```text
worker 1 mengerjakan job 1: kerjakan job ke 1
worker 2 mengerjakan job 2: kerjakan job ke 2
worker 3 mengerjakan job 3: kerjakan job ke 3
worker 1 mengerjakan job 4: kerjakan job ke 4
worker 2 mengerjakan job 5: kerjakan job ke 5
worker 3 mengerjakan job 6: kerjakan job ke 6
worker 1 mengerjakan job 7: kerjakan job ke 7
worker 2 mengerjakan job 8: kerjakan job ke 8
tugas selesai!
```

---

## 📌 Rencana Pengembangan (Roadmap)

- [ ] Implementasi algoritma Token Bucket / Leaky Bucket kustom.
- [ ] Dukungan rate limiter dinamis (per-client / per-IP).
- [ ] Integrasi middleware HTTP rate limiter (contoh: Gin, Fiber, atau `net/http`).
- [ ] Penggunaan Redis untuk distributed rate limiting.
- [ ] Unit testing dan benchmark.

---

## 📄 Lisensi

Proyek ini dibuat untuk tujuan pembelajaran dan pengembangan lebih lanjut. Bebas digunakan dan dimodifikasi.

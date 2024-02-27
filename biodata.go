package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type Teman struct {
	Absen    int
	Nama     string
	Alamat   string
	Pekerjaan string
}

// Fungsi untuk mendapatkan data teman berdasarkan absen
func getDataTeman(absen int) Teman {
	temanData := map[int]Teman{
		1: {Absen: 1, Nama: "Lee Jeno", Alamat: "Korea Selatan", Pekerjaan: "idol"},
		2: {Absen: 2, Nama: "Lee Taeyong", Alamat: "Korea Selatan", Pekerjaan: "idol"},
		3: {Absen: 3, Nama: "Song Kang", Alamat: "Korea Selatan", Pekerjaan: "aktor"},
		4: {Absen: 4, Nama: "Abimana Mahanama", Alamat: "Jakarta", Pekerjaan: "teknisi"},
		5: {Absen: 5, Nama: "Nguyen Trun", Alamat: "Vietnam", Pekerjaan: "auditor"},
		6: {Absen: 6, Nama: "Vivi Fiona", Alamat: "Jakarta", Pekerjaan: "editor"},
	}

	teman, exists := temanData[absen]
	if !exists {
		fmt.Println("Teman dengan absen", absen, "tidak ditemukan.")
		os.Exit(1)
	}

	return teman
}

// Fungsi untuk menampilkan data teman
func tampilkanDataTeman(teman Teman) {
	fmt.Println("Data Teman dengan Absen", teman.Absen)
	fmt.Println("Nama     :", teman.Nama)
	fmt.Println("Alamat   :", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <absen>")
		os.Exit(1)
	}

	// Ambil absen dari argumen command line
	absen := os.Args[1]

	// Konversi absen menjadi integer
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Absen harus berupa angka.")
		os.Exit(1)
	}

	teman := getDataTeman(absenInt)
	tampilkanDataTeman(teman)
}

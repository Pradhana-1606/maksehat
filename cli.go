package main

import "fmt"

// fungsi menampilkan menu versi cli
func cliMode() {
	var choice int

	clearConsole()

	for {
		fmt.Println("=========================================================")
		fmt.Println("               SELAMAT DATANG di makSehat")            
		fmt.Println("  Aplikasi Manajemen Kesehatan Mental - Self Assessment")
		fmt.Println("=========================================================")
		fmt.Println()
		fmt.Println("1. Tambah Data Assessment")
		fmt.Println("2. Ubah Data Assessment")
		fmt.Println("3. Hapus Data Assessment")
		fmt.Println("4. Tampilkan Data Assessment")
		fmt.Println("5. Cari Data Assessment")
		fmt.Println("6. Urutkan Data Assessment")
		fmt.Println("7. Laporan Ringkasan")
		fmt.Println("8. Keluar")
		fmt.Println()
		fmt.Println("---------------------------------------------------------")
		fmt.Print("Pilih menu [1-8]: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			// addData()
		case 2:
			// updateData()
		case 3:
			// deleteData()
		case 4:
			// showAllData()
		case 5:
			// searchData()
		case 6:
			// sortData()
		case 7:
			// showReport()
		case 8:
			fmt.Println()
			println("Program selesai, semua data yang belum disimpan telah dihapus.")
			fmt.Println()
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
			pressEnter()
		}
	}
}

// fungsi untuk menerima input string

// fungsi untuk menerima input integer
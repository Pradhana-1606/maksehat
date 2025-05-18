package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// fungsi utama mode cli
func cliMode() {
	for {
		clearConsole()

		showMenu()

		choice := intInput()

		switch choice {
		case 1:
			handleAddAssessment()
		case 2:
			// handleUpdateAssessment()
		case 3:
			// handleDeleteAssessment()
		case 4:
			// handleShowAllAssessment()
		case 5:
			// handleSearchAssessment()
		case 6:
			// handleSortAssessment()
		case 7:
			// handleShowReport()
		case 8:
			fmt.Println()
			println("Program selesai, semua data yang belum disimpan telah dihapus.")
			fmt.Println()
			return
		default:
			fmt.Println()
			fmt.Println("Pilihan tidak valid, coba lagi.")
			fmt.Println()
			pressEnter()
		}
	}
}

// fungsi handler addAssessment
func handleAddAssessment() {
	var (
		isNewUser string
		name string
	)

	clearConsole()
	showVerificationHeader()

	for {
		fmt.Print("- Apakah Anda adalah pengguna baru? (y/n): ")
		isNewUser = stringInput()
		err := yesNoValidation(isNewUser)

		if err != nil {
			fmt.Println()
			fmt.Println("  Error:", err)
			fmt.Println()
			continue
		}

		break
	}

	fmt.Println()

	for {
		fmt.Print("- Masukkan nama lengkap Anda: ")
		name = stringInput()
		err := nameInputValidation(name)

		if err != nil {
			fmt.Println()
			fmt.Println("  Error:", err)
			fmt.Println()
			continue
		}

		break
	}

	fmt.Println()
	pressEnter()

	// lanjutkan sampai pemanggilan fungsi addassessment
}

// fungsi handler updateAssessment

// fungsi handler deleteAssessment

// fungsi handler showAllAssessment

// fungsi handler searchAssessment

// fungsi handler sortAssessment

// fungsi handler showReport

// fungi untuk menampilkan daftar menu
func showMenu() {
	fmt.Println("=========================================================")
	fmt.Println("               SELAMAT DATANG di makSehat")            
	fmt.Println("  Aplikasi Manajemen Kesehatan Mental - Self Assessment")
	fmt.Println("=========================================================")
	fmt.Println()
	fmt.Println("1. Kerjakan Assessment")
	fmt.Println("2. Perbaiki Data Assessment")
	fmt.Println("3. Hapus Data Assessment")
	fmt.Println("4. Tampilkan Data Assessment")
	fmt.Println("5. Cari Data Assessment")
	fmt.Println("6. Urutkan Data Assessment")
	fmt.Println("7. Laporan Ringkasan")
	fmt.Println("8. Keluar")
	fmt.Println()
	fmt.Println("---------------------------------------------------------")
	fmt.Print("Pilih menu [1-8]: ")
}

// fungsi untuk menampilkan header verifikasi data pengguna
func showVerificationHeader() {
	fmt.Println("================================================================")
	fmt.Println("                    VERIFIKASI DATA PENGGUNA")
	fmt.Println("================================================================")
	fmt.Println()
	fmt.Println("Silakan verifikasi diri Anda sebelum mengisi kuesioner.")
	fmt.Println()
}

// fungsi untuk membersihkan terminal
func clearConsole(){
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// fungsi untuk menerima input string
func stringInput() string {
	reader := bufio.NewReader((os.Stdin))
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

// fungsi untuk menerima input integer
func intInput() int {
	source := stringInput()
	input, _ := strconv.Atoi(source)
	return input
}

// fungsi validasi yes or no
func yesNoValidation(input string) error {
	if len(input) == 0 {
		return errors.New("input tidak boleh kosong")
	}
	if input != "y"&& input != "n" || !stringInputValidation(input) {
		return errors.New("input harus y/n")
	}
	return nil
}

// fungsi untuk membaca enter
func pressEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tekan ENTER untuk melanjutkan ...")
	fmt.Print("\033[?25l")
	reader.ReadString('\n')
	fmt.Print("\033[?25h")
}
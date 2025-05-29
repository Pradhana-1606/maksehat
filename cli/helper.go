package cli

import (
	"bufio"
	"errors"
	"fmt"
	"maksehat/internal/util"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func showWelcome() {
	fmt.Println("========================================================")
	fmt.Println("               SELAMAT DATANG DI makSehat")
	fmt.Println("      Aplikasi Manajemen Kesehatan Mental Berbasis")
	fmt.Println("                    Self Assessment")
	fmt.Println("========================================================")
	fmt.Println()
	fmt.Println("1. Masuk Dengan Akun Yang Sudah Ada")
	fmt.Println("2. Daftar Akun Baru")
	fmt.Println("3. Keluar")
	fmt.Println()
	fmt.Println("--------------------------------------------------------")
	fmt.Print("Pilih menu [1-3]: ")
}

func showLoginHeader() {
	fmt.Println("========================================================")
	fmt.Println("                         MASUK")
	fmt.Println("========================================================")
}

func showRegisterHeader() {
	fmt.Println("========================================================")
	fmt.Println("                      PENDAFTARAN")
	fmt.Println("========================================================")

}

func showMenu() {
	fmt.Println("=========================================================")
	fmt.Println("               SELAMAT DATANG di makSehat")
	fmt.Println("  Aplikasi Manajemen Kesehatan Mental - Self Assessment")
	fmt.Println("=========================================================")
	fmt.Println()
	fmt.Println("1. Kerjakan Assessment")
	fmt.Println("2. Perbaiki Data Assessment")
	fmt.Println("3. Hapus Data Assessment")
	fmt.Println("4. Riwayat Assessment")
	fmt.Println("5. Cari Data Assessment")
	fmt.Println("6. Urutkan Data Assessment")
	fmt.Println("7. Laporan Ringkasan")
	fmt.Println("8. Simpan Data Assessment")
	fmt.Println("9. Edit Akun")
	fmt.Println("10. Keluar Dari Akun")
	fmt.Println("11. Keluar Dari Aplikasi")
	fmt.Println()
	fmt.Println("---------------------------------------------------------")
	fmt.Print("Pilih menu [1-11]: ")
}

func showQuestionnaireHeader() {
	fmt.Println("===========================================================================")
	fmt.Println("                    KUESIONER KESEHATAN MENTAL makSehat")
	fmt.Println("===========================================================================")
	fmt.Println()
	fmt.Println("Petunjuk pengisian kuesioner:")
	fmt.Println()
	fmt.Println("Anda akan diminta untuk menjawab 10 pertanyaan")
	fmt.Println("menggunakan skala Likert 1-5 dengan ketentuan sebagai berikut:")
	fmt.Println()
	fmt.Println("1 = Tidak Pernah")
	fmt.Println("2 = Jarang")
	fmt.Println("3 = Kadang-kadang")
	fmt.Println("4 = Sering")
	fmt.Println("5 = Selalu")
	fmt.Println()
	fmt.Println("Contoh penggunaan skala:")
	fmt.Println("Seberapa sering Anda merasa cemas akhir-akhir ini?")
	fmt.Println("Jawaban : 5")
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------")
}

func showResultHeader() {
	fmt.Println("=========================================================================")
	fmt.Println("                    HASIL ASSESSMENT KESEHATAN MENTAL")
	fmt.Println("=========================================================================")
}

func showUpdateAssessmentHeader() {
	fmt.Println("============================================================")
	fmt.Println("                  PERBAIKI DATA ASSESSMENT")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Println("Berikut data riwayat assessment yang pernah dikerjakan")
	fmt.Println()
}

func showAllAssessmentHeader() {
	fmt.Println("===========================================================")
	fmt.Println("            RIWAYAT ASSESSMENT KESEHATAN MENTAL")
	fmt.Println("===========================================================")
	fmt.Println()
	fmt.Println("Berikut data riwayat assessment yang pernah dikerjakan")
	fmt.Println()
}

func showDeleteAssessmentHeader() {
	fmt.Println("===========================================================")
	fmt.Println("                   HAPUS DATA ASSESSMENT")
	fmt.Println("===========================================================")
	fmt.Println()
	fmt.Println("Berikut data riwayat assessment yang pernah dikerjakan")
	fmt.Println()
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func stringInput() string {
	reader := bufio.NewReader((os.Stdin))
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func intInput() (int, error) {
	source := stringInput()
	input, err := strconv.Atoi(source)
	if err != nil {
		return 0, errors.New("harus angka integer")
	}
	return input, nil
}

func yesNoValidation(input string) error {
	isValid := util.StringInputValidation(input)
	if isValid != nil {
		return isValid
	}
	if input != "y" && input != "n" {
		return errors.New("harus y/n")
	}
	return nil
}

func pressEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tekan ENTER untuk melanjutkan ...")
	fmt.Print("\033[?25l")
	reader.ReadString('\n')
	fmt.Print("\033[?25h")
}

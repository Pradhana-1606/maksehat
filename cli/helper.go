package cli

import (
	"bufio"
	"errors"
	"fmt"
	"maksehat/data"
	"maksehat/internal/util"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func showMenu() {
	for _, a := range data.Assessments{
		fmt.Println(a.AssessmentID)
		fmt.Println(a.Date.Format("02-01-2006"))
		fmt.Println(a.UserID)
		fmt.Println(a.UserName)
		for _, q := range a.Answers {
			fmt.Print(q.QuestionID, ", ")
		}
		fmt.Println()
		for _, w := range a.Answers {
			fmt.Print(w.Answer, ", ")
		}
		fmt.Println()
		fmt.Println(a.TotalScore)
		fmt.Println(a.Category)
	}
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
	fmt.Println("8. Simpan Data Assessment")
	fmt.Println("9. Keluar")
	fmt.Println()
	fmt.Println("---------------------------------------------------------")
	fmt.Print("Pilih menu [1-9]: ")
}

func showVerificationHeader() {
	fmt.Println("================================================================")
	fmt.Println("                    VERIFIKASI DATA PENGGUNA")
	fmt.Println("================================================================")
	fmt.Println()
	fmt.Println("Silakan verifikasi diri Anda sebelum mengisi kuesioner.")
	fmt.Println()
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
		return 0, errors.New("harus angka")
	}
	return input, nil
}

func yesNoValidation(input string) error {
	isValid := util.StringInputValidation(input)
	if isValid != nil {
		return isValid
	}
	if input != "y"&& input != "n" {
		return errors.New("input harus y/n")
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
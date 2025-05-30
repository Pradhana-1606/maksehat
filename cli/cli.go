package cli

import (
	"fmt"
	"maksehat/data"
	"maksehat/internal/auth"
	"maksehat/internal/model"
	"maksehat/internal/service"
	"maksehat/internal/util"
	"os"
	"time"
)

// lanjut menghubungan fitur login ke main menu

func CliMode() {
	clearConsole()
	err := data.IsDBExist("data/user.json")
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		fmt.Println()
		pressEnter()
	}
	err = service.IsAdminExist()
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		fmt.Println()
		pressEnter()
	}
	for {
		clearConsole()
		showWelcome()
		choice, _ := intInput()
		switch choice {
		case 1:
			handleLogin()
		case 2:
			handleRegister()
		case 3:
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

func handleLogin() {
	clearConsole()
	showLoginHeader()
	fmt.Println()
	fmt.Print("Username : ")
	username := stringInput()
	fmt.Print("Password : ")
	password := stringInput()
	fmt.Println()
	user, err := auth.Login(username, password)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		pressEnter()
	} else {
		auth.SetActiveUser(user)
		mainMenu()
	}
}

func handleRegister() {
	var (
		name        string
		gender      string
		dateOfBirth time.Time
		username    string
		password    string
		err         error
	)
	clearConsole()
	showRegisterHeader()
	fmt.Println()
	for {
		fmt.Print("Nama Lengkap  : ")
		name = util.ToUpperCase(stringInput())
		err := util.StringInputValidation(name)
		if err != nil {
			fmt.Println()
			fmt.Println("Error: Nama", err)
			fmt.Println()
			continue
		}
		break
	}

	fmt.Println()
	fmt.Println("1. Laki-laki")
	fmt.Println("2. Perempuan")
	fmt.Println()
	for {
		fmt.Print("Jenis Kelamin [1/2]: ")
		choice, err := intInput()
		if err != nil {
			fmt.Println()
			fmt.Println("Error:", err)
			fmt.Println()
			continue
		}
		if choice < 1 || choice > 2 {
			fmt.Println()
			fmt.Println("jenis kelamin hanya ada dua, pilih salah satu")
			fmt.Println()
			continue
		}
		if choice == 1 {
			gender = "male"
		} else {
			gender = "female"
		}
		break
	}

	fmt.Println()
	for {
		fmt.Print("Tanggal Lahir DD-MM-YYYY: ")
		date := stringInput()
		dateOfBirth, err = time.Parse("02-01-2006", date)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("gunakan format DD-MM-YYYY tanpa spasi")
			fmt.Println("contoh: 06-06-2006")
			continue
		}
		break
	}

	fmt.Println()
	fmt.Println("---------------------------------------------------")
	fmt.Println()
	for {
		for {
			fmt.Print("Username : ")
			username = stringInput()
			err := auth.UsernameValidator(username)
			if err != nil {
				fmt.Println()
				fmt.Println("Error:", err)
				fmt.Println()
				continue
			}
			break
		}
		for {
			fmt.Print("Password : ")
			password = stringInput()
			err = auth.PasswordValidator(password)
			if err != nil {
				fmt.Println()
				fmt.Println("Error:", err)
				fmt.Println()
				continue
			}
			break
		}
		fmt.Println()
		err := auth.Register(name, gender, username, password, dateOfBirth)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		fmt.Println("Berhasil mendaftar!")
		fmt.Println()
		pressEnter()
		break
	}
}

func mainMenu() {
	user := auth.GetActiveUser()
	err := data.IsDBExist("data/assessment.json")
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		fmt.Println()
		pressEnter()
	}
	err = service.LoadFromDatabase()
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		fmt.Println()
		pressEnter()
	}
	for {
		clearConsole()
		showMenu()
		choice, _ := intInput()
		switch choice {
		case 1:
			handleAddAssessment(user.UserID)
		case 2:
			handleUpdateAssessment(user.UserID)
		case 3:
			handleDeleteAssessment(user.UserID)
		case 4:
			handleHistoryAssessment(user.UserID)
		case 5:
			if auth.IsAdmin() {
				handleSearchAssessment()
			} else {
				fmt.Println()
				fmt.Println("Fitur ini belum tersedia.")
				fmt.Println()
				pressEnter()
			}
		case 6:
			// handleSortAssessment()
		case 7:
			// handleShowReport()
		case 8:
			err := service.SaveToDatabase()
			if err != nil {
				fmt.Println()
				fmt.Println(err)
				fmt.Println()
				pressEnter()
			} else {
				fmt.Println()
				fmt.Println("Data assessment berhasil disimpan ke database.")
				fmt.Println()
				pressEnter()
			}
		case 9:
			// editAccount()
		case 10:
			auth.Logout()
			return
		case 11:
			fmt.Println()
			println("Program selesai, semua data yang belum disimpan telah dihapus.")
			fmt.Println()
			os.Exit(0)
		default:
			fmt.Println()
			fmt.Println("Pilihan tidak valid, coba lagi.")
			fmt.Println()
			pressEnter()
		}
	}
}

func handleAddAssessment(userID string) {
	var (
		answers    []model.Answer
		question   string
		questionID string
	)

	clearConsole()
	showQuestionnaireHeader()
	fmt.Println()

	for i := 0; i < 10; i++ {
		util.GetQuestion(10)
		questionID = data.SelectedQuestions[i].QuestionID
		question = data.SelectedQuestions[i].QuestionText
		fmt.Printf("%d. %s", i+1, question)
		fmt.Println()
		for {
			if i < 9 {
				fmt.Print("   Jawabanmu: ")
			} else {
				fmt.Print("    Jawabanmu: ")
			}
			input, err := intInput()
			if err != nil {
				if i < 9 {
					fmt.Println("   Error: Jawaban", err)
				} else {
					fmt.Println("    Error: Jawaban", err)
				}
				continue
			}
			if input < 1 || input > 5 {
				if i < 9 {
					fmt.Println("   Error: Jawaban harus diantara 1-5!")
				} else {
					fmt.Println("    Error: Jawaban harus diantara 1-5!")
				}
				continue
			}
			answers = append(answers, model.Answer{
				QuestionID: questionID,
				Answer:     input,
			})
			break
		}
		fmt.Println()
	}

	service.AddAssessment(userID, answers)

	util.ResetSelectedQuestion()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("Jawaban berhasil disimpan!")
	fmt.Println()
	pressEnter()

	clearConsole()
	showResultHeader()
	fmt.Println()
	fmt.Println("SKOR     : ", service.ScoreCalculation(answers))
	fmt.Println("KATEGORI : ", service.Categorization(service.ScoreCalculation(answers)))
	fmt.Println()
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("REKOMENDASI UMUM")
	fmt.Println()
	fmt.Println(service.Recommendation(service.ScoreCalculation(answers)))
	fmt.Println()
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println()
	pressEnter()
}

func handleHistoryAssessment(userID string) {
	clearConsole()
	showAllAssessmentHeader()
	assessments := data.Assessments

	if auth.IsAdmin() {
		if len(assessments) == 0 {
			fmt.Println("Belum ada data assessment.")
			fmt.Println()
			pressEnter()
		} else {
			fmt.Println("==========================================================================================================")
			fmt.Println("| No. | ID ASSESSMENT | TANGGAL    | ID PENGGUNA | NAMA PENGGUNA                  | SKOR | KATEGORI      |")
			fmt.Println("==========================================================================================================")
			for i := 0; i < len(assessments); i++ {
				fmt.Printf("| %3d | %10s    | %10s | %10s  | %-30s |  %3d | %-13s |",
					i+1,
					assessments[i].AssessmentID,
					assessments[i].Date.Format("02-01-2006"),
					assessments[i].UserID,
					assessments[i].Name,
					assessments[i].TotalScore,
					assessments[i].Category,
				)
				fmt.Println()
			}
			fmt.Println("==========================================================================================================")
		}
	} else {
		fmt.Println("===========================================================")
		fmt.Println("| No. | ID ASSESSMENT | TANGGAL    | SKOR | KATEGORI      |")
		fmt.Println("===========================================================")
		found := false
		count := 0
		for i := 0; i < len(assessments); i++ {
			if userID == assessments[i].UserID {
				count += 1
				fmt.Printf("| %3d | %10s    | %10s |  %3d | %-13s |",
					count,
					assessments[i].AssessmentID,
					assessments[i].Date.Format("02-01-2006"),
					assessments[i].TotalScore,
					assessments[i].Category,
				)
				fmt.Println()
				found = true
			}
		}
		fmt.Println("===========================================================")
		if !found {
			clearConsole()
			fmt.Println("Belum ada data assessment.")
			fmt.Println()
			pressEnter()
		}
	}
	fmt.Println()
	pressEnter()
}

func handleUpdateAssessment(userID string) {
	var (
		selectedAssessment string
		target             *model.Assessment
	)
	clearConsole()
	showUpdateAssessmentHeader()
	assessments := data.Assessments

	if auth.IsAdmin() {
		if len(assessments) == 0 {
			fmt.Println("Belum ada data assessment.")
			fmt.Println()
			pressEnter()
		} else {
			fmt.Println("==========================================================================================================")
			fmt.Println("| No. | ID ASSESSMENT | TANGGAL    | ID PENGGUNA | NAMA PENGGUNA                  | SKOR | KATEGORI      |")
			fmt.Println("==========================================================================================================")
			for i := 0; i < len(assessments); i++ {
				fmt.Printf("| %3d | %10s    | %10s | %10s  | %-30s |  %3d | %-13s |",
					i+1,
					assessments[i].AssessmentID,
					assessments[i].Date.Format("02-01-2006"),
					assessments[i].UserID,
					assessments[i].Name,
					assessments[i].TotalScore,
					assessments[i].Category,
				)
				fmt.Println()
			}
			fmt.Println("==========================================================================================================")
		}
	} else {
		fmt.Println("===========================================================")
		fmt.Println("| No. | ID ASSESSMENT | TANGGAL    | SKOR | KATEGORI      |")
		fmt.Println("===========================================================")
		found := false
		count := 0
		for i := 0; i < len(assessments); i++ {
			if userID == assessments[i].UserID {
				count += 1
				fmt.Printf("| %3d | %10s    | %10s |  %3d | %-13s |",
					count,
					assessments[i].AssessmentID,
					assessments[i].Date.Format("02-01-2006"),
					assessments[i].TotalScore,
					assessments[i].Category,
				)
				fmt.Println()
				found = true
			}
		}
		fmt.Println("===========================================================")
		if !found {
			clearConsole()
			fmt.Println("Belum ada data assessment.")
			fmt.Println()
			pressEnter()
		}
	}

	fmt.Println()
	for {
		fmt.Print("Masukkan ID assessment yang ingin di ubah: ")
		selectedAssessment = stringInput()
		for i := 0; i < len(assessments); i++ {
			if selectedAssessment == assessments[i].AssessmentID {
				if auth.IsAdmin() || assessments[i].UserID == userID {
					target = &assessments[i]
					break
				}
			}
		}
		if target == nil {
			fmt.Println()
			fmt.Printf("Assessment dengan ID: %s tidak ditemukan.", selectedAssessment)
			fmt.Println()
			continue
		}
		break
	}

	fmt.Println()
	fmt.Println("--- Pilih Bagian Yang Ingin Diubah ---")
	fmt.Println()
	fmt.Println("1. Tanggal Assessment")
	fmt.Println("2. Batal")
	fmt.Println()
	for {
		fmt.Print("Pilih [1-2]: ")
		selectedOperation, err := intInput()
		if err != nil {
			fmt.Println()
			fmt.Println("Error: Input", err)
			fmt.Println()
			continue
		} else {
			switch selectedOperation {
			case 1:
				fmt.Println()
				var newDate time.Time
				for {
					fmt.Print("Masukkan tanggal baru dengan format DD-MM-YYYY: ")
					newDate, err = time.Parse("02-01-2006", stringInput())
					if err != nil {
						fmt.Println("Error:", err)
						fmt.Println("gunakan format DD-MM-YYYY tanpa spasi")
						fmt.Println("contoh: 16-12-2024")
						continue
					}
					break
				}
				target.Date = newDate
				service.UpdateAssessment(target.AssessmentID, *target)
				fmt.Println()
				fmt.Println("Data berhasil diubah!")
				fmt.Println()
			case 2:
				return
			default:
				fmt.Println()
				fmt.Println("Pilihan tidak valid, coba lagi.")
				fmt.Println()
				continue
			}
		}
		break
	}
	pressEnter()
}

func handleDeleteAssessment(userID string) {
	var (
		selectedAssessment string
		target             *model.Assessment
	)
	clearConsole()
	showDeleteAssessmentHeader()
	assessments := data.Assessments

	if auth.IsAdmin() {
		if len(assessments) == 0 {
			fmt.Println("Belum ada data assessment.")
			fmt.Println()
			pressEnter()
		} else {
			fmt.Println("==========================================================================================================")
			fmt.Println("| No. | ID ASSESSMENT | TANGGAL    | ID PENGGUNA | NAMA PENGGUNA                  | SKOR | KATEGORI      |")
			fmt.Println("==========================================================================================================")
			for i := 0; i < len(assessments); i++ {
				fmt.Printf("| %3d | %10s    | %10s | %10s  | %-30s |  %3d | %-13s |",
					i+1,
					assessments[i].AssessmentID,
					assessments[i].Date.Format("02-01-2006"),
					assessments[i].UserID,
					assessments[i].Name,
					assessments[i].TotalScore,
					assessments[i].Category,
				)
				fmt.Println()
			}
			fmt.Println("==========================================================================================================")
		}
	} else {
		fmt.Println("===========================================================")
		fmt.Println("| No. | ID ASSESSMENT | TANGGAL    | SKOR | KATEGORI      |")
		fmt.Println("===========================================================")
		found := false
		count := 0
		for i := 0; i < len(assessments); i++ {
			if userID == assessments[i].UserID {
				count += 1
				fmt.Printf("| %3d | %10s    | %10s |  %3d | %-13s |",
					count,
					assessments[i].AssessmentID,
					assessments[i].Date.Format("02-01-2006"),
					assessments[i].TotalScore,
					assessments[i].Category,
				)
				fmt.Println()
				found = true
			}
		}
		fmt.Println("===========================================================")
		if !found {
			clearConsole()
			fmt.Println("Belum ada data assessment.")
			fmt.Println()
			pressEnter()
		}
	}

	fmt.Println()
	for {
		fmt.Print("Masukkan ID assessment yang ingin di ubah: ")
		selectedAssessment = stringInput()
		for i := 0; i < len(assessments); i++ {
			if selectedAssessment == assessments[i].AssessmentID {
				if auth.IsAdmin() || assessments[i].UserID == userID {
					target = &assessments[i]
					break
				}
			}
		}
		if target == nil {
			fmt.Println()
			fmt.Printf("Assessment dengan ID: %s tidak ditemukan.", selectedAssessment)
			fmt.Println()
			continue
		}
		break
	}

	fmt.Println()
	for {
		fmt.Print("Apakah Anda yakin akan menghapus data assessment tersebut? (y/n): ")
		confirm := util.ToLowerCase(stringInput())
		err := yesNoValidation(confirm)
		if err != nil {
			fmt.Println()
			fmt.Println("Input", err)
			fmt.Println()
			continue
		}
		if confirm == "y" {
			service.DeleteAssessment(target.AssessmentID)
			fmt.Println()
			fmt.Println("Assessment berhasil dihapus!")
			fmt.Println()
		} else {
			fmt.Println()
			fmt.Println("Dibatalkan!")
			fmt.Println()
		}
		break
	}

	pressEnter()
}

func handleSearchAssessment() {
	var (
		choice  int
		err     error
		results []model.Assessment
		userID string
	)
	clearConsole()
	showSearchHeader()

	found := false
	for {
		fmt.Print("Masukkan ID pengguna yang ingin dicari: ")
		userID = stringInput()
		for i := 0; i < len(data.Assessments); i++ {
			if userID == data.Assessments[i].UserID {
				found = true
				break
			}
		}
		if !found {
			fmt.Println()
			fmt.Println("Pengguna tidak ditemukan.")
			fmt.Println()
			continue
		}
		break
	}

	fmt.Println()
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Println("3. Batal")
	fmt.Println()
	for {
		fmt.Print("Pilih metode pencarian [1-3]: ")
		choice, err = intInput()
		if err != nil {
			fmt.Println()
			fmt.Println("Error: Input ", err)
			fmt.Println()
			continue
		} else {
			if choice == 1 {
				results = service.SequentialSearch(userID)
			} else if choice == 2 {
				results = service.BinarySearch(userID)
			}
		}
		break
	}

	if len(results) == 0 {
		fmt.Println()
		fmt.Println("Data tidak ditemukan.")
		fmt.Println()
		pressEnter()
	}

	fmt.Println()
	fmt.Println("Berikut data assessment yang ditemukan:")
	fmt.Println()

	fmt.Println("==========================================================================================================")
	fmt.Println("| No. | ID ASSESSMENT | TANGGAL    | ID PENGGUNA | NAMA PENGGUNA                  | SKOR | KATEGORI      |")
	fmt.Println("==========================================================================================================")
	for i := 0; i < len(results); i++ {
		fmt.Printf("| %3d | %10s    | %10s | %10s  | %-30s |  %3d | %-13s |",
			i+1,
			results[i].AssessmentID,
			results[i].Date.Format("02-01-2006"),
			results[i].UserID,
			results[i].Name,
			results[i].TotalScore,
			results[i].Category,
		)
		fmt.Println()
	}
	fmt.Println("==========================================================================================================")

	fmt.Println()
	pressEnter()
}

package cli

import (
	"fmt"
	"maksehat/data"
	"maksehat/internal/model"
	"maksehat/internal/service"
	"maksehat/internal/util"
	"strconv"
)

func CliMode() {
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

func handleAddAssessment() {
	var (
		answers   []model.Answer
		err       error
		isNewUser string
		question   string
		questionID string
		name      string
		userID    string
	)

	clearConsole()
	showVerificationHeader()

	for {
		fmt.Print("- Apakah Anda adalah pengguna baru? (y/n): ")
		isNewUser = util.ToLowerCase(stringInput())
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
		err = util.StringInputValidation(name)

		if err != nil {
			fmt.Println()
			fmt.Println("  Error:", err)
			fmt.Println()
			continue
		}

		if isNewUser == "y" {
			userID = util.GenerateUserID()
		} else {
			uname := util.ToUpperCase(name)
			userID, err = util.GetUserID(uname)
			if err != nil {
				fmt.Println()
				fmt.Println("  Error:", err)
				fmt.Println()
				continue
			}
		}

		break
	}

	fmt.Println()
	pressEnter()
	clearConsole()
	showQuestionnaireHeader()
	fmt.Println()

	for i := 0; i < 10; i++ {
		util.GetQuestion()
		questionID = data.SelectedQuestions[i].QuestionID
		question = data.SelectedQuestions[i].QuestionText
		fmt.Printf("%d. %s", i + 1, question)
		fmt.Println()
		for {
			fmt.Print("   Jawabanmu: ")
			input := intInput()
			if input < 1 || input > 5 {
				fmt.Println("   Err: Jawaban harus diantara 1-5!")
				continue
			}
			cek := strconv.Itoa(input)
			err := util.IntInputValidation(cek)
			if err != nil {
				fmt.Println("   Err: Jawaban", err)
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

	service.AddAssessment(name, userID, answers)

	util.ResetSelectedQuestion()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("Jawaban berhasil disimpan!")
	fmt.Println()
	pressEnter()
}
package cli

import (
	"fmt"
	"maksehat/data"
	"maksehat/internal/model"
	"maksehat/internal/service"
	"maksehat/internal/util"
)

func CliMode() {
	clearConsole()
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
		answers    []model.Answer
		err        error
		isNewUser  string
		question   string
		questionID string
		name       string
		userID     string
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

	failCount := 0
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
				failCount += 1
				if failCount > 2 {
					clearConsole()
					fmt.Println()
					fmt.Println("Yuh koh kelalen jenenge dewek wkwkwk")
					fmt.Println()
					pressEnter()
					return
				}
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

	service.AddAssessment(name, userID, answers)

	util.ResetSelectedQuestion()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("Jawaban berhasil disimpan!")
	fmt.Println()
	pressEnter()
	clearConsole()
	showResult()
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

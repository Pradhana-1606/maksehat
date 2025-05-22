package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// konversi ke lowercase
func toLowerCase(text string) string {
	return strings.ToLower(text)
}

// konversi ke uppercase
func toUpperCase(text string) string {
	return strings.ToUpper(text)
}

// validasi input string
func stringInputValidation(input string) bool {
	if input[0] == ' ' || input[len(input) - 1] == ' ' {
		return false
	}

	prevSpace := false
	for _, char := range input {
		if char == ' ' {
			if prevSpace {
				return false
			}
			prevSpace = true
		} else {
			if !unicode.IsLetter(char) {
				return false
			}
			prevSpace = false
		}
	}
	return true
}

// validasi input integer
func intInputValidation(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("input tidak boleh kosong")
	}
	_, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("input harus angka integer")
	}
	return nil
}

// validasi nama
func nameInputValidation(input string) error {
	if len(input) == 0 {
		return errors.New("nama tidak boleh kosong")
	}
	if !stringInputValidation(input) {
		return errors.New("nama harus huruf dan tidak memiliki spasi ganda")
	}
	return nil
}

// generate ID assessment (sementara)
func generateAssessmentID(date time.Time, score int) string {
	year := (date.Year()) % 100
	month := date.Month()
	category := 0
	if score >= 85 {
		category = 1
	} else if score >= 70 && score <= 84 {
		category = 2
	} else if score >= 55 && score <= 69 {
		category = 3
	} else if score >= 40 && score <= 54 {
		category = 4
	} else {
		category = 5
	}

	maxCount := 0
	newCount := 0
	if len(assessments) > 0 {
		currentYearMonth := fmt.Sprintf("%d%02d", year, month)
		for i := 0; i < len(assessments); i++ {
			savedYearMonth := fmt.Sprintf("%d%02d", assessments[i].date.Year(), assessments[i].date.Month())
			if currentYearMonth == savedYearMonth {
				count, _ := strconv.Atoi(assessments[i].assessmentID[4:])
				if count > maxCount {
					maxCount = count
				}
			}
		}
		newCount = maxCount + 1
	} else {
		newCount = len(assessments) + 1
	}

	id := fmt.Sprintf("A%d%02d%d%04d", year, month, category, newCount)
	return id
}

// generate ID pengguna (sementara)
func generateUserID() string {
	year := generateDate().Year() % 100
	count := rand.Intn(9999) + 1
	id := fmt.Sprintf("%d0612%04d", year, count)
	return id
}

// mengambil ID pengguna terbaru
func getLatestUID() string {
	dataCount := len(assessments)
	if dataCount == 0 {
		return "0"
	}
	latestID := assessments[0].userID
	for _, i := range assessments {
		if i.userID > latestID {
			latestID = i.userID
		}
	}
	return latestID
}

// mendapatkan ID pengguna lama
func getUserID(name string) (string, error) {
	for _, n := range assessments {
		if name == n.userName {
			return n.userID, nil
		}
	}
	return "", fmt.Errorf("pengguna dengan nama %s tidak ditemukan", name)
}

// mengecek ketersediaan ID pengguna
func isIdExist(id string) bool {
	for _, i := range assessments {
		if i.userID == id {
			return true
		}
	}
	return false
}

// generate tanggal
func generateDate() time.Time {
	year := time.Now().Year()
	month := time.Month(rand.Intn(12) + 1)
	max := 31
	switch month {
	case 4, 6, 9, 11:
		max = 30
	case 2:
		if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
			max = 29
		} else {
			max = 28
		}
	}
	day := rand.Intn(max) + 1
	Date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return Date
}

// mengambil id soal secara acak
func getQuestionsID() string {
	for {
		i := rand.Intn(len(questionBank))
		id := questionBank[i].questionID
		if !selectedQuestion[id] {
			selectedQuestion[id] = true
			return id
		}
	}
}

// mereset isi slice selectedQuestion
func resetSelectedQuestion() {
	selectedQuestion = make(map[string]bool)
}

// mengambil teks soal dari id
func getQuestionsText(id string) (string, error) {
	qid := id
	for _, inq := range questionBank {
		if inq.questionID == qid {
			return inq.questionText, nil
		}
	}
	return "", fmt.Errorf("gagal memuat soal, %s tidak ditemukan", qid)
}

// penghitungan skor
func scoreCalculation(answers []answer) int {
	totalScore := 0
	for _, s := range answers {
		totalScore += (6 - s.answer) * 2
	}
	return totalScore
}

// kategorisasi berdasarkan skor total
func categorization(score int) string {
	if score >= 85 {
		return "Stabil"
	} else if score >= 70 && score <= 84 {
		return "Cukup Stabil"
	} else if score >= 55 && score <= 69 {
		return "Tidak Stabil"
	} else if score >= 40 && score <= 54 {
		return "Depresi Ringan"
	} else {
		return "Depresi Berat"
	}
}
package util

import (
	"errors"
	"fmt"
	"maksehat/data"
	"maksehat/internal/model"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func ToLowerCase(text string) string {
	return strings.ToLower(text)
}

func ToUpperCase(text string) string {
	return strings.ToUpper(text)
}

func StringInputValidation(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("input tidak boleh kosong")
	}
	if input[0] == ' ' || input[len(input) - 1] == ' ' {
		return errors.New("tidak boleh diawali atau diakhiri dengan spasi")
	}
	runeString := []rune(input)
	prevSpace := false
	for i := 0; i < len(runeString); i++ {
		if runeString[i] == ' ' {
			if prevSpace {
				return errors.New("tidak boleh mengandung spasi ganda")
			}
			prevSpace = true
		} else {
			if !unicode.IsLetter(runeString[i]) {
				return errors.New("harus huruf")
			}
			prevSpace = false
		}
	}
	return nil
}

func IntInputValidation(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("tidak boleh kosong")
	}
	_, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("harus angka integer")
	}
	return nil
}

func GenerateAssessmentID(date time.Time, score int) string {
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
	if dataCount() > 0 {
		currentYearMonth := fmt.Sprintf("%d%02d", year, month)
		for i := 0; i < dataCount(); i++ {
			savedYearMonth := fmt.Sprintf("%d%02d", data.Assessments[i].Date.Year(), data.Assessments[i].Date.Month())
			if currentYearMonth == savedYearMonth {
				count, _ := strconv.Atoi(data.Assessments[i].AssessmentID[4:])
				if count > maxCount {
					maxCount = count
				}
			}
		}
		newCount = maxCount + 1
	} else {
		newCount = dataCount() + 1
	}

	id := fmt.Sprintf("A%d%02d%d%04d", year, month, category, newCount)
	return id
}

func GenerateUserID() string {
	year := GenerateDate().Year() % 100
	count := rand.Intn(9999) + 1
	id := fmt.Sprintf("%d0612%04d", year, count)
	return id
}

func GetUserID(name string) (string, error) {
	for i := 0; i < dataCount(); i++ {
		if name == data.Assessments[i].UserName {
			return data.Assessments[i].UserID, nil
		}
	}
	return "", errors.New("pengguna tidak ditemukan")
}

func GenerateDate() time.Time {
	year := time.Now().Year()
	month := time.Month(rand.Intn(12) + 1)
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
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
	Date := time.Date(year, month, day, hour, minute, second, 0, time.UTC)
	return Date
}

func GetQuestion() {
	for {
		i := rand.Intn(len(data.QuestionBank))
		id := data.QuestionBank[i].QuestionID
		text := data.QuestionBank[i].QuestionText
		if !isQuestionUsed(id) {
			data.SelectedQuestions = append(data.SelectedQuestions, model.Question{
				QuestionID: id,
				QuestionText: text,
			})
		}
	}
}

func isQuestionUsed(id string) bool {
	for i := 0; i < len(data.SelectedQuestions); i++ {
		if id == data.SelectedQuestions[i].QuestionID {
			return true
		}
	}
	return false
}

func ResetSelectedQuestion() {
	data.SelectedQuestions = []model.Question{}
}

func ScoreCalculation(answers []model.Answer) int {
	totalScore := 0
	for i := 0; i < len(answers); i++ {
		totalScore += (6 - answers[i].Answer) * 2
	}
	return totalScore
}

func Categorization(score int) string {
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

func dataCount() int {
	return len(data.Assessments)
}
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// fungsi konversi ke lowercase
func toLowerCase(text string) string {
	return strings.ToLower(text)
}

// fungsi konversi ke uppercase
func toUpperCase(text string) string {
	return strings.ToUpper(text)
}

// fungsi generate ID assessment (sementara)
func generateAssessmentID() string {
	year := (generateDate().Year()) % 100
	month := generateDate().Month()
	day := generateDate().Day()
	count := len(assessments) + 1
	id := fmt.Sprintf("A%d%02d%02d%03d", year, month, day, count)
	return id
}

// fungsi generate ID pengguna (sementara)
func generateUserID() string {
	year := generateDate().Year() % 100
	count := rand.Intn(9999) + 1
	id := fmt.Sprintf("%d0612%04d", year, count)
	return id
}

// fungsi generate tanggal
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

// fungsi mengambil ID pengguna terbaru
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

// fungsi mengecek ketersediaan ID pengguna
func isIdExist(id string) bool {
	for _, i := range assessments {
		if i.userID == id {
			return true
		}
	}
	return false
}

// fungsi untuk mendapatkan soal acak (getQuestions)
func getQuestions() string {
	if len(selectedQuestion) > 10 {
		selectedQuestion = make(map[string]bool)
	}
	for {
		i := rand.Intn(len(questionBank))
		id := questionBank[i].questionID
		if !selectedQuestion[id] {
			selectedQuestion[id] = true
			return id
		}
	}
}
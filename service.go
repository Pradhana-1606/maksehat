package main

import "time"

// menambahkan assessment
func addAssessment(name, userID string, questions []question, answers []answer) {
	var (
		assessmentID string
		category string
		date time.Time
		newAssessment assessment
		userName string
		totalScore int
	)

	totalScore = scoreCalculation(answers)
	date = generateDate()
	assessmentID = generateAssessmentID(date, totalScore)
	userName = toUpperCase(name)
	category = categorization(totalScore)

	newAssessment = assessment{
		assessmentID: assessmentID,
		date: date,
		userID: userID,
		userName: userName,
		questions: questions,
		answers: answers,
		totalScore: totalScore,
		category: category,
	}

	assessments = append(assessments, newAssessment)
}

// mengubah assessment

// menghapus assessment

// menampilkan assessment

// mencari assessment

// mengurutkan assessment

// menampilkan laporan
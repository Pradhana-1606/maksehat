package service

import (
	"maksehat/data"
	"maksehat/internal/model"
	"maksehat/internal/util"
	"time"
)

func AddAssessment(name, userID string, answers []model.Answer) {
	var (
		assessmentID  string
		category      string
		date          time.Time
		newAssessment model.Assessment
		userName      string
		totalScore    int
	)

	totalScore = ScoreCalculation(answers)
	date = util.GenerateDate()
	assessmentID = util.GenerateAssessmentID(date, totalScore)
	userName = util.ToUpperCase(name)
	category = Categorization(totalScore)

	newAssessment = model.Assessment{
		AssessmentID: assessmentID,
		Date:         date,
		UserID:       userID,
		UserName:     userName,
		Answers:      answers,
		TotalScore:   totalScore,
		Category:     category,
	}

	data.Assessments = append(data.Assessments, newAssessment)
}
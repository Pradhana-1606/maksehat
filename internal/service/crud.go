package service

import (
	"maksehat/data"
	"maksehat/internal/model"
	"maksehat/internal/util"
	"time"
)

func AddAssessment(userID string, answers []model.Answer) {
	var (
		assessmentID  string
		category      string
		date          time.Time
		newAssessment model.Assessment
		name      string
		totalScore    int
	)

	totalScore = ScoreCalculation(answers)
	date = util.GenerateDate()
	assessmentID = util.GenerateAssessmentID(date, totalScore)
	name, _ = GetName(userID)
	category = Categorization(totalScore)

	newAssessment = model.Assessment{
		AssessmentID: assessmentID,
		Date:         date,
		UserID:       userID,
		Name:         name,
		Answers:      answers,
		TotalScore:   totalScore,
		Category:     category,
	}

	data.Assessments = append(data.Assessments, newAssessment)
}

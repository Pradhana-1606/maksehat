package model

import "time"

type Assessment struct {
	AssessmentID string
	Date         time.Time
	UserID       string
	UserName     string
	Answers      []Answer
	TotalScore   int
	Category     string
}

type Question struct {
	QuestionID   string
	QuestionText string
}

type Answer struct {
	QuestionID string
	Answer     int
}
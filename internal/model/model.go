package model

import "time"

type Assessment struct {
	AssessmentID string    `json:"assessmentId"`
	Date         time.Time `json:"date"`
	UserID       string    `json:"userId"`
	UserName     string    `json:"userName"`
	Answers      []Answer  `json:"answers"`
	TotalScore   int       `json:"totalScore"`
	Category     string    `json:"category"`
}

type Question struct {
	QuestionID   string
	QuestionText string
}

type Answer struct {
	QuestionID string
	Answer     int
}
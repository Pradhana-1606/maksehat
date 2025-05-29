package model

import "time"

type Assessment struct {
	AssessmentID string    `json:"assessmentId"`
	Date         time.Time `json:"date"`
	UserID       string    `json:"userId"`
	Name         string    `json:"userName"`
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

type User struct {
	UserID      string
	Name        string
	Gender      string
	DateOfBirth time.Time
	Username    string
	Password    string
	IsAdmin     bool
}

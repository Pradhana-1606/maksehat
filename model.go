package main

import "time"

type assessment struct {
	assessmentID string
	date time.Time
	userID string
	userName string
	questions []question
	answers []answer
	totalScore int
	category string
}

type question struct {
	questionID string
	questionText string
}

type answer struct {
	questionID string
	answer int
}
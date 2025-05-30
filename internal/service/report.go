package service

import (
	"maksehat/internal/model"
	"sort"
	"time"
)

func GetLastFive(assessments []model.Assessment) []model.Assessment {
	sort.Slice(assessments, func(i, j int) bool {
		return assessments[i].Date.After(assessments[j].Date)
	})

	if len(assessments) > 5 {
		return assessments[:5]
	}
	return assessments
}

func GetAverageScore(assessments []model.Assessment) float64 {
	var totalScore int
	var count int

	dateNow := time.Now()
	oneMonthAgo := dateNow.AddDate(0, -1, 0)

	for i := 0; i < len(assessments); i++ {
		if assessments[i].Date.After(oneMonthAgo) {
			totalScore += assessments[i].TotalScore
			count++
		}
	}

	if count == 0 {
		return 0.0
	} else {
		return float64(totalScore) / float64(count)
	}
}

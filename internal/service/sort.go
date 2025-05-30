package service

import (
	"maksehat/data"
	"maksehat/internal/auth"
	"maksehat/internal/model"
)

func SelectionSort(userID string) []model.Assessment {
	var filtered []model.Assessment
	sorted := data.Assessments
	copy(sorted, data.Assessments)

	for i := 0; i < len(sorted)-1; i++ {
		min := i
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].TotalScore < sorted[min].TotalScore {
				min = j
			}
		}
		temp := sorted[i]
		sorted[i] = sorted[min]
		sorted[min] = temp
	}

	if !auth.IsAdmin() {
		for i := 0; i < len(sorted); i++ {
			if sorted[i].UserID == userID {
				filtered = append(filtered, sorted[i])
			}
		}
		return filtered
	}
	return sorted
}

func InsertionSort(userID string) []model.Assessment {
	var filtered []model.Assessment
	sorted := data.Assessments
	copy(sorted, data.Assessments)

	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for ; j >= 0 && sorted[j].Date.After(key.Date); j-- {
			sorted[j + 1] = sorted[j]
		}
		sorted[j + 1] = key
	}

	if !auth.IsAdmin() {
		for i := 0; i < len(sorted); i++ {
			if sorted[i].UserID == userID {
				filtered = append(filtered, sorted[i])
			}
		}
		return filtered
	}
	return sorted
}

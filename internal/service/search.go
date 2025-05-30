package service

import (
	"maksehat/data"
	"maksehat/internal/model"
	"sort"
)

func SequentialSearch(userID string) []model.Assessment {
	var results []model.Assessment
	for i := 0; i < len(data.Assessments); i++ {
		if userID == data.Assessments[i].UserID {
			results = append(results, data.Assessments[i])
		}
	}
	return results
}

func BinarySearch(userID string) []model.Assessment {
	var results []model.Assessment

	sort.Slice(data.Assessments, func(i, j int) bool {
		return data.Assessments[i].UserID < data.Assessments[j].UserID
	})
	
	low, high := 0, len(data.Assessments) - 1
	for low <= high {
		middle := (low + high) / 2
		if data.Assessments[middle].UserID == userID {
			left, right := middle, middle
			for left >= 0 && data.Assessments[left].UserID == userID {
				left--
			}
			for right < len(data.Assessments) && data.Assessments[right].UserID == userID {
				right++
			}
			for i := left + 1; i < right; i++ {
				results = append(results, data.Assessments[i])
			}
			break
		} else if data.Assessments[middle].UserID < userID {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return results
}
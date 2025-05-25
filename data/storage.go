package data

import (
	"encoding/json"
	"fmt"
	"maksehat/internal/model"
	"os"
)

func IsDBExist(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("database error: %v", err)
		}
		file, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("gagal membuat database: %v", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "\t")
		err = encoder.Encode([]any{})
		if err != nil {
			return fmt.Errorf("gagal menginisialisasi database: %v", err)
		}
	}
	return nil
}

func SaveDataAssessment(newAssessments []model.Assessment) error {
	filepath := "data/assessment.json"
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("gagal mengakses database: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(newAssessments)
	if err != nil {
		return fmt.Errorf("gagal menyimpan data: %v", err)
	}
	return nil
}

func LoadDataAssessment() ([]model.Assessment, error) {
	var oldAssessments []model.Assessment
	filepath := "data/assessment.json"

	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return oldAssessments, nil
		}
		return nil, fmt.Errorf("gagal mengakses database: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&oldAssessments)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca data: %v", err)
	}
	return oldAssessments, nil
}

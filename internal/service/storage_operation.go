package service

import (
	"errors"
	"maksehat/data"
)

func LoadFromDatabase() error {
	oldAssessments, err := data.LoadDataAssessment()
	if err != nil {
		return err
	} else {
		data.Assessments = oldAssessments
		return nil
	}
}

func SaveToDatabase() error {
	if len(data.Assessments) == 0 {
		return errors.New("belum ada data yang bisa disimpan")
	}
	err := data.SaveDataAssessment(data.Assessments)
	if err != nil {
		return err
	}
	return nil
}

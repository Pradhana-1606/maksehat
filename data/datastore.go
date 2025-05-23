package data

import "maksehat/internal/model"

var Assessments []model.Assessment

var QuestionBank = []model.Question {
	{QuestionID: "Q01", QuestionText: "Seberapa sering Anda merasa cemas atau gugup tanpa alasan yang jelas?"},
	{QuestionID: "Q02", QuestionText: "Seberapa mudah Anda tertidur di malam hari?"},
	{QuestionID: "Q03", QuestionText: "Seberapa sering Anda merasa lelah tanpa energi?"},
	{QuestionID: "Q04", QuestionText: "Seberapa sering Anda menikmati aktivitas sehari-hari?"},
	{QuestionID: "Q05", QuestionText: "Seberapa sering Anda merasa sulit berkonsentrasi?"},
	{QuestionID: "Q06", QuestionText: "Seberapa sering Anda merasa dihargai oleh orang sekitar?"},
	{QuestionID: "Q07", QuestionText: "Seberapa sering Anda merasa putus asa atau tidak ada harapan?"},
	{QuestionID: "Q08", QuestionText: "Seberapa baik nafsu makan Anda belakangan ini?"},
	{QuestionID: "Q09", QuestionText: "Seberapa sering Anda merasa puas dengan kehidupan Anda saat ini?"},
	{QuestionID: "Q10", QuestionText: "Seberapa sering Anda berpikir untuk mencari bantuan profesional?"},
}

var SelectedQuestions []model.Question
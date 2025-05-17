package main

var assessments []assessment

var questionBank = []question {
	{"Q01", "Seberapa sering Anda merasa cemas atau gugup tanpa alasan yang jelas?"},
	{"Q02", "Seberapa mudah Anda tertidur di malam hari?"},
	{"Q03", "Seberapa sering Anda merasa lelah tanpa energi?"},
	{"Q04", "Seberapa sering Anda menikmati aktivitas sehari-hari?"},
	{"Q05", "Seberapa sering Anda merasa sulit berkonsentrasi?"},
	{"Q06", "Seberapa sering Anda merasa dihargai oleh orang sekitar?"},
	{"Q07", "Seberapa sering Anda merasa putus asa atau tidak ada harapan?"},
	{"Q08", "Seberapa baik nafsu makan Anda belakangan ini?"},
	{"Q09", "Seberapa sering Anda merasa puas dengan kehidupan Anda saat ini?"},
	{"Q10", "Seberapa sering Anda berpikir untuk mencari bantuan profesional?"},
}

var selectedQuestion = make(map[string]bool)
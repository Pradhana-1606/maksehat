package service

import "maksehat/internal/model"

func ScoreCalculation(answers []model.Answer) int {
	totalScore := 0
	for i := 0; i < len(answers); i++ {
		totalScore += (6 - answers[i].Answer) * 2
	}
	return totalScore
}

func Categorization(score int) string {
	if score >= 85 {
		return "Stabil"
	} else if score >= 70 && score <= 84 {
		return "Cukup Stabil"
	} else if score >= 55 && score <= 69 {
		return "Tidak Stabil"
	} else if score >= 40 && score <= 54 {
		return "Depresi Ringan"
	} else {
		return "Depresi Berat"
	}
}

func Recommendation(score int) string {
	if score >= 85 {
		return "Kondisi Anda sangat baik. Pertahankan rutinitas sehat, jaga pola tidur,\n\ndan lakukanlah aktifitas positif yang membuat Anda\n\nmerasa tenang dan bahagia."
	} else if score >= 70 && score <= 84 {
		return "Kondisi Anda sudah cukup baik, namun ada beberapa tanda stres ringan.\n\nPastikan Anda punya waktu untuk relaksasi dan\n\nkomunikasikan dengan orang-orang terdekat jika diperlukan."
	} else if score >= 55 && score <= 69 {
		return "Anda menunjukkan gejala ketidakstabilan emosional. Kurangi beban kerja,\n\nambil jeda istirahat, dan pertimbangkan untuk berkonsultasi\n\ndengan profesional jika kondisi ini berlanjut."
	} else if score >= 40 && score <= 54 {
		return "Terdapat tanda-tanda depresi ringan dalam jawaban Anda.\n\nCobalah mencari dukungan sosial, mulai kegiatan kecil yang menyenangkan,\n\ndan pertimbangkan untuk berkonsultasi dengan psikolog."
	} else {
		return "Gejala yang muncul menunjukkan bahwa Anda mengalami depresi berat.\n\nKami sangat menyarankan Anda untuk segera menghubungi profesional\n\nkesehatan mental dan mendapatkan bantuan yang tepat."
	}
}
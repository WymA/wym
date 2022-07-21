package internal

import "time"

func GetDailyFileName() string {
	return "../assets/words/" + time.Now().AddDate(0, 0, 1).Format("20060102") + ".json"
}

type Word struct {
	Index      int    `json:"index"`
	Word       string `json:"word"`
	Definition string `json:"def"`
}

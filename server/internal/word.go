package internal

import "time"

func GetTodaysDateWithFormat(format string) string {
	return time.Now().AddDate(0, 0, 1).Format(format)
}

func GetTodaysDate() string {
	return time.Now().AddDate(0, 0, 1).Format("20060102")
}

func GetDailyFileName() string {
	return "../assets/words/" + GetTodaysDate() + ".json"
}

type Word struct {
	Index      int    `json:"index"`
	Word       string `json:"word"`
	Definition string `json:"def"`
}

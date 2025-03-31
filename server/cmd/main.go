package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/WymA/wym/server/internal"
)

const EVERYDAY_WORDS_NUM = 12

func main() {
	dictionaryJson, err := os.Open(internal.PUBLIC_PATH + "assets/websters/dictionary.json")

	if err != nil {
		log.Fatalln(err)
	}

	defer dictionaryJson.Close()

	dictionaryStat, err := dictionaryJson.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Dictionary file size: %d bytes", dictionaryStat.Size())

	words, err := internal.Decode2WordsInLimit(dictionaryJson)
	if err != nil {
		log.Fatalln(err)
	}

	todayWords := make([]internal.Word, 0)
	totalWordsNum := len(words)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < EVERYDAY_WORDS_NUM; i++ {

		pick := r.Intn(totalWordsNum)
		log.Printf("Word: %+v", words[pick])
		todayWords = append(todayWords, words[pick])
	}

	totaysJson, _ := json.MarshalIndent(todayWords, "", " ")

	_ = os.WriteFile(internal.GetDailyFileName(), totaysJson, 0644)

	currentDir, err := os.Getwd()
	internal.CheckErr(err)
	log.Println("Current working directory:", currentDir)

	internal.GenDailyWordHtmlFromJson(internal.GetDailyFileName(), internal.HISTORY_PATH)
	internal.CreateHistoryHtmlAndSitemap(internal.HISTORY_PATH)
}

package internal

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/template"
)

func GenDailyWordHtmlFromJson(dailyNewWordJsonFilePath string, historyPath string) {
	template, err := template.New("newWordItem").Parse(WORD_INDEX_HEADER + DIV_NEW_WORD_ITEM_TEMPLATE + WORD_INDEX_FOOTER)
	CheckErr(err)

	dailyNewWordJsonFile, err := os.Open(dailyNewWordJsonFilePath)
	CheckErr(err)
	defer dailyNewWordJsonFile.Close()

	todayWords := make([]Word, 0)

	todayWords, err = ReadLimitedJsonFile2Words(dailyNewWordJsonFile)
	CheckErr(err)

	data := struct {
		TodayDate string
		Items     []Word
	}{
		TodayDate: GetTodaysDateWithFormat("02/01/2006"),
		Items:     todayWords,
	}

	historyFile := fmt.Sprintf(historyPath+"word-%s.html", GetTodaysDate())
	log.Println(historyFile)

	wordHistoryHtml, err := os.Create(historyFile)
	CheckErr(err)

	indexFile := fmt.Sprintf(PUBLIC_PATH + "index.html")
	log.Println(indexFile)

	indexHtml, err := os.Create(indexFile)
	CheckErr(err)

	err = template.Execute(wordHistoryHtml, data)
	CheckErr(err)
	err = template.Execute(indexHtml, data)
	CheckErr(err)
}

func GetHistoryFilenameFromFolder(historyPath string) []string {
	files, err := os.ReadDir(historyPath)
	CheckErr(err)

	fileNames := make([]string, 0)
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}

func GetDateFromFilename(filename string) string {
	// Extract date from format "word-20250319.html"
	if len(filename) < 15 {
		return ""
	}

	dateStr := filename[5:13] // Extract "20250319"

	// Format as "2025-03-19" for better readability
	if len(dateStr) == 8 {
		return dateStr[:4] + dateStr[4:6] + dateStr[6:8]
	}

	return dateStr
}

func GetFilenameNoExt(filename string) string {
	// GetFilenameNoExt returns the filename without its extension
	// Example: "word-20250319.html" -> "word-20250319"
	if len(filename) > 4 && filename[len(filename)-4:] == ".htm" || filename[len(filename)-5:] == ".html" {
		if filename[len(filename)-4:] == ".htm" {
			return filename[:len(filename)-4]
		}
		return filename[:len(filename)-5]
	}
	return filename
}

func CreateHistoryHtmlAndSitemap(historyPath string) {
	template, err := template.New("allHistoryItems").Parse(HISTORY_HEADER + DIV_ALL_HISTORY_ITEMS + WORD_INDEX_FOOTER)
	CheckErr(err)

	filenames := GetHistoryFilenameFromFolder(historyPath)
	historyItems := make([]History, 0)

	for i, filename := range filenames {
		if filename == "index.html" || filename == "history.html" || len(filename) < 18 {
			continue
		}

		// Extract filename without extension
		filenameNoExt := GetFilenameNoExt(filename)
		date := GetDateFromFilename(filename)

		historyItems = append(historyItems, History{
			Index:    i + 1,
			Filename: filenameNoExt,
			Date:     date,
		})
	}

	data := struct {
		Items []History
	}{
		Items: historyItems,
	}

	// Sort history items from latest date to oldest
	sort.Slice(historyItems, func(i, j int) bool {
		// Compare dates in reverse order (latest first)
		return historyItems[i].Date > historyItems[j].Date
	})

	historyHtml, err := os.Create(historyPath + "history.html")
	CheckErr(err)

	err = template.Execute(historyHtml, data)
	CheckErr(err)

	sitemap := GenerateSitemapFromHistory(historyItems)

	sitemapFile, err := os.Create(historyPath + "sitemap.xml")
	CheckErr(err)

	err = sitemap.Generate(sitemapFile)
	CheckErr(err)

}

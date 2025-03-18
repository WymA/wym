package internal

import (
	"fmt"
	"os"
	"text/template"
)

const WORD_INDEX_HEADER string = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="vocabulary builder, 12 new words everyday">
    <meta name="author" content="matthias2wym, wym, matthias, matthias.eth, WymA, vocabulary builder">
    <link rel="icon" href="assets/img/favicon.ico">

    <title>Wym's home | 12 new words everyday</title>

    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap/4.6.1/css/bootstrap.min.css" integrity="sha512-T584yQ/tdRR5QwOpfvDfVQUidzfgc2339Lc8uBDtcp/wYu80d7jwBgAxbyMh0a9YM9F8N3tdErpFI8iaGx6x5g==" crossorigin="anonymous" referrerpolicy="no-referrer" />

    <link  rel="stylesheet" href="https://matthias2wym.com/assets/css/cards.css"></style>
    <style>
        body {
        padding-top: 5rem;
        }
        .starter-template {
        padding: 3rem 1.5rem;
        text-align: center;
        }
    </style>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-3Z4DCRGB21"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'G-3Z4DCRGB21');
    </script>
    {{range .Items}}
    <meta name="{{.Word}}" content="{{.Definition}}">
    {{end}}
  </head>

  <body>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
      <a class="navbar-brand" href="#">Wym</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item active">
            <a class="nav-link" href="https://matthias2wym.com/index.html">Home <span class="sr-only">(current)</span></a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="https://github.com/WymA">github</a>
          </li>
        </ul>
        <!-- <form class="form-inline my-2 my-lg-0">
          <input class="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search">
          <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
        </form> -->
      </div>
    </nav>

    <main role="main" class="container">
`

const DIV_NEW_WORD_ITEM_TEMPLATE = `
      <div class="starter-template">
        <h1>12 new words everyday</h1>
        <h3>{{.TodayDate}}</h2>
      </div>
      <div class="container">
        <div id="new-words" class="row">
        {{range .Items}}
          <div id="{{.Index}}" class="col-sm-4">
            <div class="card card-flip h-100">
              <div class="card-front text-white bg-dark">
                <div class="card-body">
                  <i class="fa fa-search fa-5x float-right"></i>
                  <h3 class="card-title">#{{.Index}}</h3>
                  <p class="card-text">{{.Word}}</p>
                </div>
              </div>
              <div class="card-back bg-white">
                <div class="card-body">
                  <h3 class="card-title">{{.Word}}</h3>
                  <p class="card-text">{{.Definition}}</p>
                </div>
              </div>
            </div>
          </div>
          {{else}}{{end}}
        </div>
      </div>
`

const HISTORY_HEADER string = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="vocabulary builder, 12 new words everyday">
    <meta name="author" content="matthias2wym, wym, matthias, matthias.eth, WymA, vocabulary builder">
    <link rel="icon" href="assets/img/favicon.ico">

    <title>Wym's home | 12 new words everyday</title>

    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap/4.6.1/css/bootstrap.min.css" integrity="sha512-T584yQ/tdRR5QwOpfvDfVQUidzfgc2339Lc8uBDtcp/wYu80d7jwBgAxbyMh0a9YM9F8N3tdErpFI8iaGx6x5g==" crossorigin="anonymous" referrerpolicy="no-referrer" />

    <link  rel="stylesheet" href="https://matthias2wym.com/assets/css/cards.css"></style>
    <style>
        body {
        padding-top: 5rem;
        }
        .starter-template {
        padding: 3rem 1.5rem;
        text-align: center;
        }
    </style>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-3Z4DCRGB21"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'G-3Z4DCRGB21');
    </script>
  </head>

  <body>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
      <a class="navbar-brand" href="#">Wym</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item active">
            <a class="nav-link" href="https://matthias2wym.com/index.html">Home <span class="sr-only">(current)</span></a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="https://github.com/WymA">github</a>
          </li>
        </ul>
        <!-- <form class="form-inline my-2 my-lg-0">
          <input class="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search">
          <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
        </form> -->
      </div>
    </nav>

    <main role="main" class="container">
`

const DIV_ALL_HISTORY_ITEMS = `
      <div class="starter-template">
        <h1>History</h1>
      </div>
      <div class="container">
        {{range .Items}}
        <div id="{{.Index}}" class="col-sm-4">
          <a href="https://matthias2wym.com/history/{{.Filename}}.html">{{.Date}}</a>
        </div>
        {{else}}{{end}}
      </div>
`

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

	wordHistoryHtml, err := os.Create(fmt.Sprintf(historyPath+"word-%s.html", GetTodaysDate()))
	CheckErr(err)

	indexHtml, err := os.Create(fmt.Sprintf(historyPath + "../index.html"))
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

func CreateHistoryHtmlFile(historyPath string) {
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

	historyHtml, err := os.Create(historyPath + "history.html")
	CheckErr(err)

	err = template.Execute(historyHtml, data)
	CheckErr(err)
}

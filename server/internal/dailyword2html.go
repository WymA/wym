package internal

import (
	"fmt"
	"os"
	"text/template"
)

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

func GenDailyWordHtmlFromJson(dailyNewWordJsonFilePath string, publicPath string) {
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

	wordHistoryHtml, err := os.Create(fmt.Sprintf(publicPath+"word-%s.html", GetTodaysDate()))
	CheckErr(err)

	indexHtml, err := os.Create(fmt.Sprintf(publicPath+"../index.html", GetTodaysDate()))
	CheckErr(err)

	err = template.Execute(wordHistoryHtml, data)
	CheckErr(err)
	err = template.Execute(indexHtml, data)
	CheckErr(err)
}

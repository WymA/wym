package internal

import (
	"os"
	"text/template"
)

const DIV_NEW_WORD_ITEM_TEMPLATE = `{{range .Items}}<div id="{{.Index}}" class="col-sm-4">
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
													<h3 class="card-title">Definition</h3>
													<p class="card-text">{{.Definition}}</p>
												</div>
											</div>
										</div>
									</div>
									{{else}}{{end}}`

func GenDailyWordHtmlFromJson(dailyNewWordJsonFilePath string) {
	template, err := template.New("newWordItem").Parse(DIV_NEW_WORD_ITEM_TEMPLATE)
	CheckErr(err)

	dailyNewWordJsonFile, err := os.Open(dailyNewWordJsonFilePath)
	CheckErr(err)
	defer dailyNewWordJsonFile.Close()

	todayWords := make([]Word, 0)

	todayWords, err = ReadLimitedJsonFile2Words(dailyNewWordJsonFile)
	CheckErr(err)

	data := struct {
		Items []Word
	}{
		Items: todayWords,
	}

	err = template.Execute(os.Stdout, data)
	CheckErr(err)
}

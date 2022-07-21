package internal

import "testing"

func TestGenDailyWordHtmlFromJson(t *testing.T) {
	type args struct {
		dailyNewWordJsonFilePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "path",
			args: args{
				dailyNewWordJsonFilePath: "../../assets/words/20220722.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenDailyWordHtmlFromJson(tt.args.dailyNewWordJsonFilePath)
		})
	}
}

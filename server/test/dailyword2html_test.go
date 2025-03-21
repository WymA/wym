package test

import (
	"testing"

	"github.com/WymA/wym/server/internal"
)

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
				dailyNewWordJsonFilePath: "../" + internal.PUBLIC_PATH + "assets/words/20220722.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			internal.GenDailyWordHtmlFromJson(tt.args.dailyNewWordJsonFilePath, "../../public/")
		})
	}
}
func TestGetDateFromFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "normal case",
			filename: "word-20250319.html",
			want:     "20250319",
		},
		{
			name:     "too short",
			filename: "short.html",
			want:     "",
		},
		{
			name:     "different date",
			filename: "word-20220101.html",
			want:     "20220101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.GetDateFromFilename(tt.filename); got != tt.want {
				t.Errorf("GetDateFromFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFilenameNoExt(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "html extension",
			filename: "word-20250319.html",
			want:     "word-20250319",
		},
		{
			name:     "htm extension",
			filename: "word-20250319.htm",
			want:     "word-20250319",
		},
		{
			name:     "no extension",
			filename: "word-20250319",
			want:     "word-20250319",
		},
		{
			name:     "other extension",
			filename: "word-20250319.json",
			want:     "word-20250319.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.GetFilenameNoExt(tt.filename); got != tt.want {
				t.Errorf("GetFilenameNoExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

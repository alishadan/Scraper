package save_string

import (
	//"database/sql"
	"encoding/json"
	"os"
	"scraper/src/core/text_type"
	"strings"
)

func Save_string_file(fileName string, content [][]string) bool {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, row := range content {
		file.WriteString(strings.Join(row, ",") + "\n")
	}

	return true

}

func Save_articles(fileName string, articles []text_type.Article) bool {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, article := range articles {
		if err := encoder.Encode(article); err != nil {
			panic(err)
			return false
		}
	}

	return true
}

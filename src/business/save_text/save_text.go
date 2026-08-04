package save_text

import (
	"scraper/src/core/text_type"
	"scraper/src/data/save_string"
)

func Save_text(filename string, data [][3]string) int {
	var articles []text_type.Article
	for index := range data {
		article := text_type.Article{
			Title: data[index][0],
			URL:   data[index][1],
			Image: data[index][2],
		}

		articles = append(articles, article)

	}
	result := save_string.Save_articles(filename, articles)
	return result
}

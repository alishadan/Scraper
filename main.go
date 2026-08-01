// main.go
package main

import (
	"fmt"
	"scraper/src/business/check_requirements"
	"scraper/src/business/get_text"
	"scraper/src/business/save_text"
	"scraper/src/data/read_input"
)

func main() {
	site, success := read_input.ReadInputFile() // Make sure function is exported
	if !success {
		fmt.Println("Failed to read input file")
		return
	}

	fmt.Println("START SCRAPING")
	articles := get_text.Scrape(site.URL, site.Tags)
	path := "data"
	filename := "data/data.json"
	check_requirements.Check_file(path, filename)
	if len(articles) == 0 {
		fmt.Println("do not find any data")
		return
	}
	result := save_text.Save_text(filename, articles)
	if result {
		fmt.Println("Data saved as JSON successfully!")
		return
	} else {
		fmt.Println("Failed to save data as JSON")
		return
	}

}

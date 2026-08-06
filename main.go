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
	fmt.Printf("start app \n")

	//step 1: read input file
	site, success := read_input.ReadInputFile()
	if success != 0 {
		fmt.Println("Failed to read input file")
		return
	}

	fmt.Println("START SCRAPING")

	//step 2:
	articles := get_text.Scrape(site.URL, site.Tags) // [][3]string
	path := "data"
	filename := "data/data.json"
	check_requirements.Check_file(path, filename)
	if len(articles) == 0 {
		fmt.Println("do not find any data")
		return
	}
	result := save_text.Save_text(filename, articles)
	println(result)
	if result == 0 {
		fmt.Println("Data saved as JSON successfully!")
		return
	} else {
		fmt.Println("Failed to save data as JSON")
		return
	}

}

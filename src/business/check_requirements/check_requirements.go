package check_requirements

import (
	"scraper/src/data/check_folder"
)

func Check_file(path string, filename string) bool {
	result := check_folder.Check_folder(path, filename)
	return result
}

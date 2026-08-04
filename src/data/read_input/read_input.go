package read_input

import (
	"encoding/json"
	"os"
	"scraper/src/core/text_type"
)

// Export the function (capitalize first letter)
func ReadInputFile() (text_type.InputData, int) {
	var inputData text_type.InputData
	inputPath := "input.json"

	// Check if file exists
	_, err := os.Stat(inputPath)
	if os.IsNotExist(err) {
		// input file is not exist
		return inputData, 1
	}

	// Read the file (using os.ReadFile - modern approach)
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		//Error reading input file
		return inputData, 2
	}

	// Unmarshal JSON into struct
	if err := json.Unmarshal(bytes, &inputData); err != nil {
		//Error parsing JSON
		return inputData, 3
	}

	return inputData, 0
}

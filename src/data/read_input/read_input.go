// scraper/src/data/read_input/read_input.go
package read_input

import (
	"encoding/json"
	"fmt"
	"os"
	"scraper/src/core/text_type"
)

// Export the function (capitalize first letter)
func ReadInputFile() (text_type.InputData, bool) {
	var inputData text_type.InputData
	inputPath := "input.json"

	// Check if file exists
	_, err := os.Stat(inputPath)
	if os.IsNotExist(err) {
		fmt.Printf("Error: input file '%s' does not exist\n", inputPath)
		return inputData, false
	}

	// Read the file (using os.ReadFile - modern approach)
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return inputData, false
	}

	// Unmarshal JSON into struct
	if err := json.Unmarshal(bytes, &inputData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return inputData, false
	}

	return inputData, true
}

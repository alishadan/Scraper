package check_folder

import (
	"errors"
	"log"
	"os"
)

func Check_folder(path string, filename string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatal(err)
			return false
		}
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
			return false
		}
		defer f.Close()
		return true

	} else {
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
				return false
			}
			defer f.Close()
		}
		return true
	}

}

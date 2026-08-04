package check_folder

import (
	"errors"
	"os"
)

func Check_folder(path string, filename string) int {
	//check and create folder
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			//erorr in make directory in path
			return 1
		}
		f, err := os.Create(filename)
		if err != nil {
			//error in make filename in path
			return 12
		}
		defer f.Close()
		return 0

	} else {
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(filename)
			if err != nil {
				//error in creating file
				return 2
			}
			defer f.Close()
		}
		return 0
	}

}

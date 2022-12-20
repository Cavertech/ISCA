// This Package performace file check for hashing, Line of Code in the project and file type

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FileType() {
	// FileType fuction get the path of the file and also the type of the file.
	// return:  with path and file type
	files, err := os.ReadDir("/home/yellowhoodie/Desktop/ISCA/core/scan_engine /scan_dir")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		file_ext := file.Name()
		fmt.Println(filepath.Ext(file_ext))
		// array for list of file ext
		strArray1 := [3]string{"Japan", "Australia", "Germany"}
		fmt.Println(strArray1)
	}

}

func LineOfCode() {
	// LineOfCode gives lines of code that is within a file
	// para

}

func main() {
	FileType()

}

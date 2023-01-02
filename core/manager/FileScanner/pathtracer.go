// Package tracksdown files absolute path with filtering for file types
// Returns:= file in /ISCA/core/scan_engine /scan_dir
package filescanner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type ProjectInfo struct {
	ProjectName string
	ScanDir     string
	TempDIr     string
}

func PathFinder() string {
	// Get the Absolute path of the defualt scan dir
	ScanDir, erro := filepath.Abs("./core/scan_engine /scan_dir")
	if erro != nil {
		log.Fatal(erro)
	}

	var FilePaths []string                  //collects and stores the path in a slice
	extentions := []string{".conf", ".txt"} // list of supported file type of vailidation

	// travel across the dir from the root to mapout the dir structure
	err := filepath.Walk(ScanDir, func(path string, info os.FileInfo, err error) error {
		//handle err func
		if err != nil {
			log.Fatal(err)
			return nil
		}

		// Loop will vaildate the file type and  add the output path into the slice FilePath
		for _, ext := range extentions {
			// check if for Dir ==true and file extention
			if !info.IsDir() && filepath.Ext(path) == ext {
				FilePaths = append(FilePaths, path)
			}
		}
		return nil
	})

	// handle err for err:= filepath.Walk
	if err != nil {
		log.Fatal(err)
	}
	var data string
	for _, file := range FilePaths {
		data = file
	}
	return data
}

// Create & write to a file.
func WriteFilePath(lines []string) {
	PathList, error := os.CreateTemp()
	if error != nil {
		log.Fatal(error)
		return
	}

}

func (pi ProjectInfo) getInfo(data string) {
	pi.ProjectName = data
	pi.ScanDir = "./core/scan_engine /scan_dir/"
	pi.TempDIr = fmt.Sprintf("./core/scan_engine /scan_dir/Reports/%s", pi.ProjectName)
}

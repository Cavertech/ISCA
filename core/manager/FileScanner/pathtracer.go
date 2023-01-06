// Package tracksdown files absolute path with filtering for file types
// Returns:= file in /ISCA/core/scan_engine /scan_dir
package filescanner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	ProjectName string
	scandir     string = "./core/scan_engine/scan_dir"
	reportdir   string = "./core/scan_engine/"
)

func PathFinder() {
	// Get the Absolute path of the defualt scan dir
	ScanDir, erro := filepath.Abs(scandir)
	if erro != nil {
		log.Fatal(erro)
	}

	var FilePaths []string               //collects and stores the path in a slice
	extentions := []string{".go", ".tf"} // list of supported file type of vailidation

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
	for _, file := range FilePaths {
		fmt.Println(file)

	}
	WriteFilePath(FilePaths, "FilePaths")

}

// Create & write to a file.
// parr slice string := filepaths
// parr string := Report patters for generating tempalets
func WriteFilePath(lines []string, ReportPattern string) {
	ReportDirPath, erro := filepath.Abs(reportdir)
	if erro != nil {
		log.Fatal(erro)
	}
	reportdir := fmt.Sprintf("%s/%s/Reports", ReportDirPath, ProjectName)
	pattern := fmt.Sprintf("%s-%s", ProjectName, ReportPattern)
	PathList, error := os.CreateTemp(reportdir, pattern)
	if error != nil {
		log.Fatal(error)
		return
	}
	defer PathList.Close()

	for _, val := range lines {
		_, writeError := fmt.Fprintln(PathList, val)
		if writeError != nil {
			log.Fatal(writeError)
			return
		}
	}

}

// Create a Dir Structure of the projects. The project name is passed as a parameter
func CreateDirStr(DirName string) string {
	root, erro := filepath.Abs(reportdir)
	if erro != nil {
		log.Fatal(erro)
	}
	dirCreate := fmt.Sprintf("%s/%s/Reports", root, ProjectName)
	fmt.Println(dirCreate)
	if err := os.MkdirAll(dirCreate, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	return "pass"
}

// Start the package
func StartFileScanner() {
	//Start by creating Dir Stuct
	CreateDirStr(ProjectName)
	fmt.Println("Completed creating Dir structure")
	fmt.Println("Started Pathfinder")
	PathFinder()
	fmt.Println("Scan compeleted")

}

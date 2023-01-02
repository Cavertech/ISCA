package main

import (
	"fmt"

	filescanner "github.com/Cavertech/ISCA/core/manager/FileScanner"
	"github.com/Cavertech/ISCA/core/manager/welcome"
)

func main() {
	fmt.Println(welcome.WelcomeMesaage())
	fmt.Println(filescanner.PathFinder())
}

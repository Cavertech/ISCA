package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorReset := "\033[0m"

	var fileName string
	var GitUrl string
	fmt.Println("Project Name:")
	fmt.Scanln(&fileName)
	fmt.Println("Enter Repository url:")
	fmt.Scanln(&GitUrl)
	_, err := git.PlainClone(fileName, false, &git.CloneOptions{
		URL:      GitUrl,
		Progress: os.Stdout,
	})
	// https://github.com/ybkuroki/go-webapp-sample.git"
	// fmt.Print(err)
	if err != nil {
		fmt.Println(string(colorRed), err, colorReset)
	} else {
		fmt.Println(string(colorGreen), "Cloned", colorReset)
	}

}

package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	//Colors for output
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorReset := "\033[0m"

	var fileName string

	var opt int

	fmt.Println("Project Name:")
	fmt.Scanln(&fileName)
	fmt.Println("Please select an option from below \n 1. Plain clone \n 2. Clone with Personal token")
	fmt.Scanln(&opt)
	switch {
	case opt == 1:
		var GitUrl string
		fmt.Println("Enter Repository url:")
		fmt.Scanln(&GitUrl)
		_, err := git.PlainClone(fileName, false, &git.CloneOptions{
			URL:      GitUrl,
			Progress: os.Stdout,
		})

		if err != nil {
			fmt.Println(string(colorRed), err, colorReset)
		} else {
			fmt.Println(string(colorGreen), "Successfully Cloned", colorReset)
		}
	case opt == 2:
		var userName string
		var key string
		var repo string
		fmt.Println("Enter username:")
		fmt.Scanln(&userName)
		fmt.Println("Enter Access token:")
		fmt.Scanln(&key)
		fmt.Println("Enter Repository url:")
		fmt.Scanln(&repo)
		_, err := git.PlainClone(fileName, false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: userName,

				Password: key,
			},
			URL: repo,

			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Println(string(colorRed), err, colorReset)
		} else {
			fmt.Println(string(colorGreen), "Successfully Cloned", colorReset)
		}
	}

}

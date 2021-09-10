package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	fileList := os.Args[1]
	clonePath := os.Args[2]

	fmt.Println(fileList)
	fmt.Println(clonePath)

	// Open the file.
	f, err := os.Open(fileList)

	if err != nil {
		fmt.Printf("file err: %v", err)
		os.Exit(-1)
	}

	gitRepo := make([]string, 0)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them
	for scanner.Scan() {
		repo := scanner.Text()
		fmt.Println(repo)
		gitRepo = append(gitRepo, repo)
	}

	for _, repo := range gitRepo {
		cmd := exec.Command("git", "clone", repo, clonePath)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		//Print the output
		fmt.Printf("%v is cloned", repo)
	}

}

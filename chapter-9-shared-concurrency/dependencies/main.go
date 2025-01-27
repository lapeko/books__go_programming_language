package main

import (
	"github.com/lapeko/books__go_programming_language/chapter-9-shared-concurrency/dependencies/utils"
	"log"
	"os"
)

func main() {
	envDeps, err := utils.GetAllEnvDeps()
	if err != nil {
		log.Fatalln(err)
	}
	wholeDepsList, err := utils.GetDeps(envDeps, os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(wholeDepsList)
}

package parser

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Read5eFile() {
	file, err := os.Open("assets/dnd5e.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Cannot Find File: ", file.Name())
		}
	}(file)

	//messages := make([]string, 0)

	fileBytes, err := os.ReadFile("assets/dnd5e.log")
	if err != nil {
		panic("Oh Shit the files Broke")
	}
	fileString := string(fileBytes[:])
	messages := strings.Split(fileString, "---------------------------")

	for _, message := range messages {
		if strings.Contains(message, "Aine Vicis") {
			fmt.Printf("=========================")
			fmt.Println(message)
		}
	}
}

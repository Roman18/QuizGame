package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
	"errors"
)

func handleError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "[-] %s\n", err.Error())
		os.Exit(1)
	}
}

func usage(args []string){
	if len(args) < 2{
		fmt.Fprintf(os.Stderr, "Usage: %s <file.csv>\n", path.Base(args[0]))
		os.Exit(1)
	}
}

func readCSVFile(file *os.File) Questions{
	questions := Questions{}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	handleError(err)

	for index, record := range records{
		questions = append(questions, Question{Id: uint(index + 1), Test: record[0], Answer: record[1]})
	}
	return questions

}

func quiz(questions Questions, interval uint) (uint, error) {
	var score uint

	timer := time.NewTimer(time.Duration(interval) * time.Second)
	for _, question := range questions{
		fmt.Printf("#%d %s\n>>> ", question.Id, question.Test)

		answer := make(chan string)
		anyError := make(chan error)
		go func(){
			scanner := bufio.NewScanner(os.Stdin)
			if ok := scanner.Scan(); !ok{
				anyError <- errors.New("Error during reading user input")
			}
			answer <- strings.TrimSpace(scanner.Text())
		}()

		select{
			case <- timer.C:
				fmt.Println("\n\tTime elapsed")
				return score, nil
			case err := <- anyError:
				return 0, err
			case yourAnswer := <- answer:
				if yourAnswer == question.Answer{score++}
		}
	}

	return score, nil
}


func main(){
	args := os.Args
	usage(args)
	fileName := args[1]

	file, err := os.Open(fileName)

	handleError(err)

	defer file.Close()

	questions := readCSVFile(file)
	var interval uint = 10

	fmt.Printf("File %s was loaded\nYou have %d seconds\n", fileName, interval)

	score, err := quiz(questions, interval)

	fmt.Printf("\nYour score: %d\n", score)

}

// struct to store single question
type Question struct{
	Id uint
	Test string
	Answer string
}


// aliace type to store bunch of questions
type Questions []Question



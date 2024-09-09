package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strings"
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


func main(){
	args := os.Args
	usage(args)
	fileName := args[1]

	file, err := os.Open(fileName)

	handleError(err)

	defer file.Close()

	fmt.Printf("File %s was opened\n", fileName)

	questions := readCSVFile(file)

}

// struct to store single question
type Question struct{
	Id uint
	Test string
	Answer string
}


// aliace type to store bunch of questions
type Questions []Question



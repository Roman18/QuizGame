package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
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

func main(){
	args := os.Args
	usage(args)
	fileName := args[1]

	file, err := os.Open(fileName)

	handleError(err)

	defer file.Close()

	fmt.Printf("File %s was opened\n", fileName)

	reader := csv.NewReader(file)
	
	records, err := reader.ReadAll()

	handleError(err)

	for _, record := range records{
		fmt.Println(record)
	}
}




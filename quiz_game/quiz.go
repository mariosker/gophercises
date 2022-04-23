package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csv_filename := flag.String("csv", "problems.csv", "A csv file containing the problems. Format: question, answer")
	flag.Parse()

	lines, err := ReadCsv(*csv_filename)

	if err != nil {
		exit((fmt.Sprintf("Failed to open %s\n", *csv_filename)), err)
	}

	correct := 0
	total := len(lines)

	var input string

	for index, line := range lines {
		data := Problem{
			Question: line[0],
			Answer:   line[1],
		}
		fmt.Println(data.Question)

		_, err = fmt.Scanln(&input)
		if err != nil {
			exit("Could not get input\n", err)
		}

		if runtime.GOOS == "windows" {
			input = strings.TrimRight(input, "\r\n")
		} else {
			input = strings.TrimRight(input, "\n")
		}

		if strings.Compare(input, data.Answer) == 0 {
			correct += 1
		}
		total = index + 1
	}

	fmt.Printf("%d correct answers out of %d", correct, total)
}

func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
		return [][]string{}, err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		log.Fatal(err)
		return [][]string{}, err
	}

	return lines, nil
}

func exit(msg string, err error) {
	log.Fatal(err)
	fmt.Println(msg)
	os.Exit(1)
}

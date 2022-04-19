package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type CsvLine struct {
	Question string
	Answer   string
}

func main() {
	lines, err := ReadCsv("problems.csv")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	correct := 0
	total := len(lines)

	var input string

	for index, line := range lines {
		data := CsvLine{
			Question: line[0],
			Answer:   line[1],
		}
		fmt.Println(data.Question)

		_, err = fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
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

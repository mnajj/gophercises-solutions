package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	csvFlag := flag.String("csv", "./ex1/problems.csv", "A CSV file in the format of \"question,answer\"")
	timeFlag := flag.String("time", "30", "A time for your test to finish within")
	flag.Parse()
	records := readCsvFile(*csvFlag)
	dur, err := strconv.Atoi(*timeFlag)
	if err != nil {
		log.Fatal("error while parsing time duration flag")
	}
	done := make(chan struct{})
	res := make(chan int)
	go takeTest(records, done, res)
	var score int
OUT:
	for {
		select {
		case score = <-res:
		case <-done:
			break OUT
		case <-time.After(time.Duration(dur) * time.Second):
			fmt.Println("Time up")
			break OUT
		}
	}
	fmt.Printf("Your scored %d out of %d\n", score, len(records))
}

func takeTest(records [][]string, done chan<- struct{}, res chan<- int) {
	var score int
	var ans string
	for i, l := range records {
		fmt.Printf("Problem #%d: %s\n", i+1, l[0])
		_, err := fmt.Scanf("%s\n", &ans)
		if err != nil {
			log.Fatal("Unable to read input file")
		}
		if ans == strings.TrimSpace(l[1]) {
			fmt.Println("Correct!")
			score++
			res <- score
			continue
		}
		fmt.Println("Wrong!")
	}
	close(done)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file"+filePath, err)
	}
	defer f.Close()
	cvsReader := csv.NewReader(f)
	records, err := cvsReader.ReadAll()
	if err != nil {
		log.Fatal("unable to parse file as CSV for "+filePath, err)
	}
	return records
}

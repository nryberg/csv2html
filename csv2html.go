package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    csv2html inputFile -in inputFile -out outputFile ...\n")
		flag.PrintDefaults()
	}
	var inputFile, outputFile string
	var cssStyleClass string

	var inHandle *os.File
	var err error

	flag.StringVar(&inputFile, "in", "", "Specify a source file")
	flag.StringVar(&outputFile, "out", "", "Specify a output file")
	flag.StringVar(&cssStyleClass, "style", "", "Specify a CSS class style")
	flag.Parse()

	// First check for stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 { // data is being piped to stdin
		inHandle = os.Stdin
	} else {
		if flag.NFlag() == 0 { // No flags, no stdin, spit out usage and quit
			fmt.Println("no args?")
			flag.Usage()
			os.Exit(1)
		} else {

			if inHandle, err = os.Open(inputFile); err != nil { // regular file input
				log.Fatal(err)
			}
		}
	}

	reader := bufio.NewReader(inHandle)
	check(err)
	r := csv.NewReader(reader)
	rowNumber := 1
	var working string
	if len(cssStyleClass) > 0 {
		working += "<table class='" + cssStyleClass + "'>\n"
	} else {
		working += "<table>\n"
	}

	var record []string
	for {
		record, err = r.Read()
		//check(err)
		if rowNumber == 1 {
			working += " <tr>\n"
			for _, v := range record {
				working += "  <th>" + v + "</th>\n"
			}
			working += " </tr>\n"
		} else {
			working += " <tr>\n"
			for _, v := range record {
				working += "  <td>" + v + "</td>\n"
			}
			working += " </tr>\n"
		}
		rowNumber++

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(record)
	}
	working += "</table>\n"

	if outputFile != "" { // if we have an output file then write it out
		out, err := os.Create(outputFile)
		check(err)
		w := bufio.NewWriter(out)
		_, err = w.WriteString(working)
		check(err)
		w.Flush()
	} else { // else send it to stdout
		fmt.Println(working)
	}
}

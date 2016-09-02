package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"flag"
	// "path"
	// "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//TODO: 			 Add markdown format
//TODO Add a help message for no args
//TODO add a parsing for std in?

func main() {
	var file_in, file_out string
	var style_class string


	flag.StringVar(&file_in,"in", "base_fruit.csv","Specify a source file")
	flag.StringVar(&file_out,"out", "test.html","Specify a output file")
	flag.StringVar(&style_class,"style", "", "Specify a CSS class style")
	flag.Parse()

	f, err := os.Open(file_in)
	check(err)
	r := csv.NewReader(bufio.NewReader(f))
	row_num := 1
	var working string
	if len(style_class) > 0 {
		working += "<table class='" + style_class + "'>\n"
	} else {
		working += "<table>\n"
	}
	for {
		record, err := r.Read()
		if row_num == 1 {
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
		row_num += 1

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(record)
	}
	working += "</table>\n"

	out, err := os.Create(file_out)
	check(err)
	w := bufio.NewWriter(out)
	_, err = w.WriteString(working)
	check(err)
	w.Flush()
	fmt.Println(working)
}

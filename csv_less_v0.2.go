/*
Author: Tyler Sammons
Date: 9/5/2020

Just a bit of code to test printing a preview organized CSV files in the terminal
*/
package main

import (
    "fmt"
    "io"
    "os"
    "encoding/csv"
    "log"
    "text/tabwriter"
    "strings"

    "github.com/fatih/color"
)

func main() {

    // filename is second commandline arg
    fname := os.Args[1]

    // formatting options
    min_cell_width := 20
    tab_width := 4
    cell_padding := 2

    // list to store headings
    var headings []string
    // create column heading to column data map
    column_data := make(map[string]string)

    // open the file
    csvfile, err := os.Open(fname)
    if err != nil {
        log.Fatalln("Couldn't open the csv file", err)
    }

    // parse the file
    r := csv.NewReader(csvfile)

    // first record is headings
    var first_record bool = true

    // iterate through the records
    for {
        // read each record from csv
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        // use tabwriter to align columns
        w := new(tabwriter.Writer)
        w.Init(os.Stdout, min_cell_width, tab_width, cell_padding, ' ', 0)

        // column position
        var pos int = 0

        if first_record { // create headings list and display fancy headings

            first_record = false // prevent rerun
            var line string

            for _, record := range record {
                headings = append(headings, record)
                line += "|" + record + "|\t"
            }
            line += "\n"
            line += strings.Repeat("-", len(record)*min_cell_width)
            line += "\n"
            color.Set(color.FgRed, color.Bold)
            fmt.Fprintf(w, line)
            w.Flush()
            color.Unset()
        } else {

            var line string

            for _, record := range record { // map heading to data and display
                column_data[headings[pos]] += record + "\n"
                pos += 1
                line += "|" + record + "\t"
            }
            line += "\n"
            fmt.Fprintf(w, line)
            w.Flush()
        }
    }
    // check mapping
    for i := 0; i < len(headings); i++ {
        fmt.Printf("\n\nHeading: %s\nData:\n%s", headings[i], column_data[headings[i]])
    }
}
/* use map to list heading and X number of lines of data
    include up arrow and down arrow to scroll data *see less
    -- reprint heading at top when scrolling
    include configuration for tabwriter in commandline
    include number of lines to display in commandline
*/

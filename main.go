package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-sql-injector <url>")
		return
	}

	url := os.Args[1]
	tables := detectTables(url)
	columns := detectColumns(url, tables)

	fmt.Println("Detected columns:")
	for table, cols := range columns {
		fmt.Printf("Table: %s\n", table)
		for _, col := range cols {
			fmt.Printf(" - %s\n", col)
		}
	}
}

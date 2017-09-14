//go 1.8
package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"sort"
)

type keyValue struct {
	Key   string
	Value int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func countIpSort (accLog []string) []keyValue {
	maxIpUsage := make(map[string]int)
	for _, line := range accLog {
		if line != "" && (strings.Count(line, " ") > 1) {
			splitLine := strings.Split(line, " ")
			maxIpUsage[splitLine[1]]++
		}
	}
	var ipSlice []keyValue
	for key, value := range maxIpUsage {
		ipSlice = append(ipSlice, keyValue{key, value})
	}

	sort.Slice(ipSlice, func(i, j int)bool {
		return ipSlice[i].Value > ipSlice[j].Value
	})
	return ipSlice
}

func main() {
	filename := os.Args[1]
	maxcount := 50
	accLog, err := readLines(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Невозможно прочитать файл: %v: %v\n", filename, err)
		os.Exit(1)
	}
	ipcount := countIpSort(accLog)

	for i:= 0; i < maxcount; i++ {
		fmt.Printf("%v requests from %s\n", ipcount[i].Value, ipcount[i].Key)
	}
}

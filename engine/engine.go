package engine

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Process(inputFile string) (string, error) {
	if strings.HasSuffix(inputFile, ".csv") {
		return processCsv(inputFile)
	} else if strings.HasSuffix(inputFile, ".prn") {
		return "dummy", nil
	} else {
		parts := strings.Split(inputFile, ".")
		extension := parts[len(parts)-1]
		return "", fmt.Errorf("unrecognized file extension %s", extension)
	}
}

func processCsv(inputFile string) (string, error) {
	header, content, err := getFileContent(inputFile)
	if err != nil {
		return "", err
	}
	headerSlice := splitStringBy(header, ",")
	contentSlice := parseCsvContent(content)

	return createHtmlTable(headerSlice, contentSlice), nil
}

func getFileContent(inputFile string) (string, []string, error) {
	lines, err := readLines(inputFile)
	if err != nil {
		return "", nil, fmt.Errorf("could not oppen file %s", inputFile)
	}
	header, content := lines[0], lines[1:]
	return header, content, nil
}

func parseCsvContent(content []string) [][]string {
	var lines [][]string
	for _, line := range content {
		var data []string
		re := regexp.MustCompile(`"(.*?)"`)
		match := re.FindString(line)
		lineWithoutName := strings.Replace(line, match+",", "", -1)
		name := strings.ReplaceAll(match, "\"", "")
		columnData := strings.Split(lineWithoutName, ",")
		data = append(data, name)
		data = append(data, columnData...)
		lines = append(lines, data)
	}
	return lines
}

func createHtmlTable(headers []string, content [][]string) string {
	tableOpeningTag := "<table>"
	tableClosingTag := "</table>"
	theadOpeningTag := "<thead>"
	theadClosingTag := "</thead>"
	tbodyOpeningTag := "<tbody>"
	tbodyClosingTag := "</tbody>"
	trOpeningTag := "<tr>"
	trClosingTag := "</tr>"
	thOpeningTag := "<th>"
	thClosingTag := "</th>"
	tdOpeningTag := "<td>"
	tdClosingTag := "</td>"

	var table string
	table += tableOpeningTag
	table += theadOpeningTag
	table += trOpeningTag
	for _, header := range headers {
		table += thOpeningTag
		table += header
		table += thClosingTag
	}
	table += theadClosingTag
	table += tbodyOpeningTag
	for _, line := range content {
		table += trOpeningTag
		for _, item := range line {
			table += tdOpeningTag
			table += item
			table += tdClosingTag
		}
		table += trClosingTag
	}
	table += tbodyClosingTag
	table += tableClosingTag

	return table
}

func splitStringBy(str string, separator string) []string {
	return strings.Split(str, separator)
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

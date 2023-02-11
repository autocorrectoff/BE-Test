package engine

import (
	"bufio"
	"fmt"
	"github.com/autocorrectoff/BE-Test/utils"
	"os"
	"regexp"
	"strings"
)

func Process(inputFile string) (string, error) {
	if strings.HasSuffix(inputFile, ".csv") {
		return processCsv(inputFile)
	} else if strings.HasSuffix(inputFile, ".prn") {
		return processPrn(inputFile)
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

func processPrn(inputFile string) (string, error) {
	header, content, err := getFileContent(inputFile)
	if err != nil {
		return "", err
	}

	delimiter := "|"
	headerSlice := splitHeaders(header, delimiter)
	contentSlice := parsePrnContent(content, len(headerSlice), delimiter)
	return createHtmlTable(headerSlice, contentSlice), nil
}

/**
*	1. Replace all consecutive whitespaces with delimiter
*	2. Replace last whitespace in the string with delimiter
*	3. Check if split line parts are equal to the number of headers
*	4. If not, check if there's more than one part containing numerics
*	5. If yes, then it's probably address and postcode. Insert delimiter between tokens containing numerics
**/
func parsePrnContent(content []string, columnCount int, delimiter string) [][]string {
	var lines [][]string
	var separatedData []string
	for _, line := range content {
		lineWithSeparator := replaceWhiteExtraWhitespacesWithDelimiter(line, delimiter)
		lineWithSeparator = replaceLastWhitespacesWithDelimiterUsingRegex(lineWithSeparator, delimiter)
		separatedData = append(separatedData, lineWithSeparator)
	}

	for _, line := range separatedData {
		var parts []string
		actualCount := strings.Count(line, delimiter)
		if actualCount != columnCount {
			parts = splitStringBy(line, delimiter)
			longest := utils.LongestString(parts)
			longestStringParts := strings.Split(longest, " ")
			var partsContainingNumericsIndexes []int
			for i, part := range longestStringParts {
				res := utils.ContainsNumerics(part)
				if res {
					partsContainingNumericsIndexes = append(partsContainingNumericsIndexes, i)
				}
			}
			if len(partsContainingNumericsIndexes) >= 2 {
				sliceStart := longestStringParts[:partsContainingNumericsIndexes[0]+1]
				sliceStartStr := strings.Join(sliceStart, " ")
				line = strings.Replace(line, sliceStartStr, sliceStartStr+delimiter, -1)
				parts = splitStringBy(line, delimiter)
			}
		} else {
			parts = splitStringBy(line, delimiter)
		}

		lines = append(lines, parts)
	}
	return lines
}

func replaceWhiteExtraWhitespacesWithDelimiter(str string, delimiter string) string {
	re := regexp.MustCompile(`\s{2,}`)
	return re.ReplaceAllString(str, delimiter)
}

func getFileContent(inputFile string) (string, []string, error) {
	lines, err := readLines(inputFile)
	if err != nil {
		return "", nil, fmt.Errorf("could not oppen file %s", inputFile)
	}
	header, content := lines[0], lines[1:]
	return header, content, nil
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

func splitHeaders(str string, separator string) []string {
	str = replaceWhitespacesWithDelimiter(str, separator)
	str = strings.Replace(str, "Credit|Limit", "Credit Limit", 1)
	return strings.Split(str, separator)
}

func replaceWhitespacesWithDelimiter(str string, delimiter string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(str, delimiter)
}

func replaceLastWhitespacesWithDelimiter(str string, separator rune) string {
	index := strings.LastIndex(str, " ")
	return utils.ReplaceAtIndex(str, separator, index)
}

func replaceLastWhitespacesWithDelimiterUsingRegex(str string, separator string) string {
	re := regexp.MustCompile(`\s(\S*)$`)
	lastWhitespace := strings.Join([]string{separator, "$1"}, "")
	return re.ReplaceAllString(str, lastWhitespace)
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

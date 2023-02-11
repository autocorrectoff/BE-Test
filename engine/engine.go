package engine

import (
	"fmt"
	"strings"
)

func Process(inputFile string) (string, error) {
	if strings.HasSuffix(inputFile, ".csv") {
		return "dummy", nil
	} else if strings.HasSuffix(inputFile, ".prn") {
		return "dummy", nil
	} else {
		parts := strings.Split(inputFile, ".")
		extension := parts[len(parts)-1]
		return "", fmt.Errorf("unrecognized file extension %s", extension)
	}
}

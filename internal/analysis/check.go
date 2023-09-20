package analysis

import (
	"fmt"
	"strings"
)

func CheckError(v string, i, j int, headers map[string]string, records [][]string, errors map[string][]string, valuesFound map[string]bool) {
	//check if the value is empty based on the header
	if v == "" && (headers[records[0][j]] == "required" || headers[records[0][j]] == "required_unique") {
		errors[fmt.Sprintf("%d%d", i, j)] = append(errors[fmt.Sprintf("%d%d", i, j)], "Missing value")
	}

	//add the value to the map
	if headers[records[0][j]] == "required_unique" {
		if valuesFound[strings.ToLower(strings.Replace(v, " ", "", -1))] {
			errors[fmt.Sprintf("%d%d", i, j)] = append(errors[fmt.Sprintf("%d%d", i, j)], v)
		} else {
			valuesFound[strings.ToLower(strings.Replace(v, " ", "", -1))] = true
		}
	}
}

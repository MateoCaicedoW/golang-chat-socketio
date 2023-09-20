package csv

import (
	"encoding/csv"
	"io"
	"main/internal/analysis"
)

func Process(file io.Reader, headers map[string]string) ([][]string, map[string][]string, error) {
	valuesFound := map[string]bool{}
	errors := map[string][]string{}
	csvFile := csv.NewReader(file)
	records, err := csvFile.ReadAll()
	if err != nil {
		return [][]string{}, errors, err
	}

	for i, record := range records {
		for j, v := range record {
			analysis.CheckError(v, i, j, headers, records, errors, valuesFound)
		}
	}

	return records, errors, nil
}

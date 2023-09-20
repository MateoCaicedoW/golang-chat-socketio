package excel

import (
	"io"
	"main/internal/analysis"

	"github.com/xuri/excelize/v2"
)

func Process(file io.Reader, headers map[string]string) ([][]string, map[string][]string, error) {
	valuesFound := map[string]bool{}
	errors := map[string][]string{}
	excelizeFile, err := excelize.OpenReader(file, excelize.Options{})
	if err != nil {
		return [][]string{}, errors, err
	}

	var matrix [][]string
	for _, sheet := range excelizeFile.GetSheetMap() {
		rows, err := excelizeFile.GetRows(sheet)
		if err != nil {
			return [][]string{}, errors, err
		}
		for i, row := range rows {
			var rowSlice []string
			for j, cell := range row {
				analysis.CheckError(cell, i, j, headers, rows, errors, valuesFound)
				rowSlice = append(rowSlice, cell)
			}
			matrix = append(matrix, rowSlice)
		}
	}

	return matrix, errors, nil
}

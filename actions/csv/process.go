package csv

import (
	"encoding/csv"
	"fmt"
	"main/render"
	"net/http"
	"path"

	"github.com/tealeg/xlsx"
)

func Process(w http.ResponseWriter, r *http.Request) {
	// decoder := form.NewDecoder()
	//parse multipart form
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//retrieve file from posted form-data
	file, handler, err := r.FormFile("File")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	//this is to save the error and the position where it happened
	errorss := []string{}

	if path.Ext(handler.Filename) == ".csv" {

		// add this csv data to a matrix
		// and pass it to the template
		// to be rendered in a table
		csvFile := csv.NewReader(file)
		records, err := csvFile.ReadAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i, v := range records {
			for j, v := range v {
				if v == "" {
					records[i][j] = "empty"
					errorss = append(errorss, fmt.Sprintf("row %d, column %d is empty", i+1, j+1))
				}
			}
		}

		for _, v := range errorss {
			fmt.Println(v)
		}

		render.SetData("matrix", records)
		render.RenderWithLayout(w, "/csv/process.html", "auth.html")
		return
	}

	xlsx, err := xlsx.OpenReaderAt(file, handler.Size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// add this xlsx data to a matrix
	// and pass it to the template
	// to be rendered in a table
	var matrix [][]string
	for _, sheet := range xlsx.Sheets {
		for rowIndex, row := range sheet.Rows {
			var rowSlice []string
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				if text == "" {
					text = "empty"
					errorss = append(errorss, fmt.Sprintf("row %d, column %d is empty", rowIndex+1, cellIndex+1))
				}
				rowSlice = append(rowSlice, text)
			}
			matrix = append(matrix, rowSlice)
		}
	}

	for _, v := range errorss {
		fmt.Println(v)
	}

	// var wg sync.WaitGroup
	// matrixChannel := make(chan [][]string, len(xlsx.Sheets))

	// for _, sheet := range xlsx.Sheets {
	// 	wg.Add(1)
	// 	go processSheet(sheet, matrixChannel, &wg)
	// }

	// go func() {
	// 	wg.Wait()
	// 	close(matrixChannel)
	// }()

	// var finalMatrix [][]string
	// for matrix := range matrixChannel {
	// 	finalMatrix = append(finalMatrix, matrix...)
	// }

	render.SetData("matrix", matrix)
	render.RenderWithLayout(w, "/csv/process.html", "auth.html")
}

// func processSheet(sheet *xlsx.Sheet, matrixChannel chan [][]string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	var matrix [][]string
// 	for _, row := range sheet.Rows {
// 		var rowSlice []string
// 		for _, cell := range row.Cells {
// 			text := cell.String()
// 			if text == "" {
// 				text = "empty"
// 			}
// 			rowSlice = append(rowSlice, text)
// 		}
// 		matrix = append(matrix, rowSlice)
// 	}
// 	matrixChannel <- matrix
// }

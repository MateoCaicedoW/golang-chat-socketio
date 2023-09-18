package csv

import (
	"encoding/csv"
	"io"
	"main/render"
	"net/http"
	"os"
	"path"
	"sync"

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

	excelFile, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//copy the uploaded file to the destination file
	if _, err := io.Copy(excelFile, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if path.Ext(handler.Filename) == ".csv" {
		// read the csv file
		a, err := os.Open(handler.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer a.Close()
		csvFile := csv.NewReader(a)
		records, err := csvFile.ReadAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, v := range records {
			for _, v2 := range v {
				if v2 == "" {
					v2 = "empty"
				}
			}
		}

		_ = os.Remove(handler.Filename)
		render.SetData("matrix", records)
		render.RenderWithLayout(w, "/csv/process.html", "auth.html")
		return
	}

	// add this csv data to a matrix
	// and pass it to the template
	// to be rendered in a table

	xlsx, err := xlsx.OpenFile(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// add this xlsx data to a matrix
	// and pass it to the template
	// to be rendered in a table
	var wg sync.WaitGroup
	matrixChannel := make(chan [][]string, len(xlsx.Sheets))

	for _, sheet := range xlsx.Sheets {
		wg.Add(1)
		go processSheet(sheet, matrixChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(matrixChannel)
	}()

	var finalMatrix [][]string
	for matrix := range matrixChannel {
		finalMatrix = append(finalMatrix, matrix...)
	}

	err = os.Remove(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.SetData("matrix", finalMatrix)
	render.RenderWithLayout(w, "/csv/process.html", "auth.html")
}

func processSheet(sheet *xlsx.Sheet, matrixChannel chan [][]string, wg *sync.WaitGroup) {
	defer wg.Done()

	var matrix [][]string
	for _, row := range sheet.Rows {
		var rowSlice []string
		for _, cell := range row.Cells {
			text := cell.String()
			if text == "" {
				text = "empty"
			}
			rowSlice = append(rowSlice, text)
		}
		matrix = append(matrix, rowSlice)
	}
	matrixChannel <- matrix
}

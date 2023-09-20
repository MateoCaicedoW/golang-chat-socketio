package csv

import (
	"main/internal/csv"
	"main/internal/excel"
	"main/render"
	"net/http"
	"path"
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

	headers := map[string]string{
		"user_ID":          "not_required",
		"User Email":       "required_unique",
		"birthday":         "not_required",
		"Employee ID":      "required",
		"Contract ID":      "required",
		"Security ID":      "required",
		"Insurance Number": "required",
	}

	if path.Ext(handler.Filename) == ".csv" {
		records, errors, err := csv.Process(file, headers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.SetData("errors", errors)
		render.SetData("matrix", records)
		render.RenderWithLayout(w, "/csv/process.html", "auth.html")
		return
	}

	records, errors, err := excel.Process(file, headers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("errors", errors)
	render.SetData("matrix", records)
	render.RenderWithLayout(w, "/csv/process.html", "auth.html")
}

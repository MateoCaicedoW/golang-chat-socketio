package csv

import (
	"main/render"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {

	render.RenderWithLayout(w, "/csv/show.html", "a.html")
}

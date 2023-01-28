package handler

import (
	"net/http"

	"github.com/msound/todo/pkg/view"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "index.html")
}

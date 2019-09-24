package function

import (
	"net/http"
)

//Handle handles the function call to function
func Handle(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect") != ""
	if redirect {
		http.Redirect(w, r, r.FormValue("redirect"), http.StatusTemporaryRedirect)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

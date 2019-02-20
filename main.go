package main

import (
	"net/http"

	"github.com/rheeger/portfolio_checker/portfolio"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/portfolio", portfolio.Index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/portfolio/show", portfolio.Show)
	http.HandleFunc("/portfolio/create", portfolio.Create)
	http.HandleFunc("/portfolio/create/process", portfolio.CreateProcess)
	http.HandleFunc("/portfolio/update", portfolio.Update)
	http.HandleFunc("/portfolio/update/process", portfolio.UpdateProcess)
	http.HandleFunc("/portfolio/delete/process", portfolio.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/portfolio", http.StatusSeeOther)
}

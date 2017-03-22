package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	/*
		Mendeklarasikan folder di mana file static tersimpan di dalam server
		Di sini nama foldernya ialah public
	*/
	fs := http.FileServer(http.Dir("public"))
	/* Mendeklarasikan url yang digunakan untuk mengakses file static nya. */
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("view", "layout.html")
	fp := filepath.Join("view", filepath.Clean(r.URL.Path))
	log.Println(lp)
	log.Println(fp)
	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

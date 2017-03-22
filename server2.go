package main

import "fmt"
import "html/template"
import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Nama":  "Subhan Nizar",
			"Pesan": "Selamat Bersenang Senang dengan golang",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		t.Execute(w, data)
	})

	fmt.Println("web server bisa di akses di http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}

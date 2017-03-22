package main
import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello Web App")
	})
	fmt.Println("web server bisa di akses di http://localhost:8000/")
	http.ListenAndServe(":8000",nil)
}

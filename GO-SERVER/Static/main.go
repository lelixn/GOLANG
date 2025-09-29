package main
import (
	"fmt"
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs) //send that to file server(fs)
	http.HandleFunc("/form", formHandler)
	 //send that to formHandler

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server at port 9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println(err)
	}
	
	// fmt.Println("server started at localhost:9000")
	// http.ListenAndServe(":9000", nil)
}
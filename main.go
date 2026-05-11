package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello from server")
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	fmt.Println(text)
	fmt.Println(banner)

	// fmt.Fprintf(w, "Hello from server")
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	ascii := GenerateAscii(text, banner)
	tmpl.Execute(w, ascii)

}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

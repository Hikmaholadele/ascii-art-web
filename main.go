package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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

func GenerateAscii(text, banner string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\\n", "\n")

	data, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return "Error reading banner file\n"
	}

	fileLines := strings.Split(string(data), "\n")
	lines := strings.Split(text, "\n")

	result := ""

	for _, line := range lines {

		// handle empty line from \n\n
		if line == "" {
			result += "\n"
			continue
		}

		// print ascii art normally
		for row := 1; row <= 8; row++ {

			for _, ch := range line {

				index := (int(ch) - 32) * 9

				result += fileLines[index+row]
			}

			result += "\n"
		}
	}

	return result
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server running on http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

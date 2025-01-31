package main

import (
	"fmt"
	"net/http"
	"os"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the index.html file from the templates directory
	file, err := os.Open("static/templates/index.html")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Error reading file information", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, "index.html", fileInfo.ModTime(), file)
}

func main() {
	// Route for serving the main HTML file
	http.HandleFunc("/", htmlHandler)

	// Route for serving static files (CSS, JS, Images)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

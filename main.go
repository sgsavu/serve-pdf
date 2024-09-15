package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var file string

func main() {
	port := flag.String("port", "8080", "Port to start the server on")
	flag.Parse()
	file = flag.Arg(0)

	if file == "" {
		log.Fatal("file needs to be provided")
	}

	http.HandleFunc("/", servePDF)

	fmt.Printf("Server is running on http://localhost:%s\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func servePDF(w http.ResponseWriter, r *http.Request) {
	pdfFile, err := os.Open(file)
	if err != nil {
		http.Error(w, "File not found.", 404)
		fmt.Println(err)
		return
	}
	defer pdfFile.Close()

	w.Header().Set("Content-Type", "application/pdf")

	http.ServeFile(w, r, file)
}

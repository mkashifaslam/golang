package main

import (
	"fmt"
	"github.com/mkashifaslam/golang/navigator/city"
	"github.com/mkashifaslam/golang/navigator/csv"
	"log"
	"net/http"
	"os"
)

func RunServer() {
	fileSystem := http.FileServer(http.Dir("data"))

	// serve static files in a directory
	http.Handle("/", fileSystem)

	http.HandleFunc("/csv", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		csvData, err := os.ReadFile(csvFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(csvData))
	})

	http.HandleFunc("/cities", func(w http.ResponseWriter, r *http.Request) {
		csvData, err := csv.LoadData(csvFile, static)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var cities string
		for i, cty := range city.GetList(csvData) {
			cities += fmt.Sprintf("<strong style='font-size:24px; font-family:Arial'>%d: %s</strong><br>", i+1, cty.Name)
		}
		fmt.Fprintf(w, "%s", cities)
	})

	// serve api with handle
	http.Handle("/api/ping", func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World")
		}
	}())

	// serve api with handler
	http.HandleFunc("/api/handler", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "func handler")
	})

	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}

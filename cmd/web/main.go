package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	PORT string = ":8080"
)

func main() {
	if len(os.Args) > 1 {
		PORT = ":" + os.Args[1]
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("..", "..", "ui", "static")))))
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/create", tableHandler)
	http.HandleFunc("/add", addTableHandler)
	http.HandleFunc("/delete", deleteFromTableHandler)
	http.HandleFunc("/get", getTableHandler)
	http.HandleFunc("/update", updateTableHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"web/model"
	// "../../internal/database/model"
)

type DataMain struct {
	Title       string
	Header      string
	Description string
	Data        []string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	data := DataMain{
		Title:       "Main page",
		Header:      "Welcome",
		Description: "Simple website in golang language",
		Data: []string{
			"Golang",
			"Html",
			"Css",
			"Sql",
		},
	}
	files := []string{
		filepath.Join("..", "..", "ui", "html", "layout", "layout.html"),
		filepath.Join("..", "..", "ui", "html", "index.html"),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func tableHandler(w http.ResponseWriter, r *http.Request) {
	model.CreateTable("users", []string{"field1", "field2", "field3"})
	fmt.Fprint(w, "Table \"users\" was created")
}

func addTableHandler(w http.ResponseWriter, r *http.Request) {
	model.AddToDB("users", []string{"field1", "field2", "field3"}, []string{"name", "surname", "description"})
	fmt.Fprint(w, "Information has been added to the database")
}

func deleteFromTableHandler(w http.ResponseWriter, r *http.Request) {
	model.DeleteFromDB("users", 2)
	fmt.Fprint(w, "Information has been deleted")
}

func getTableHandler(w http.ResponseWriter, r *http.Request) {
	info := model.GetFromDB("users")
	for _, i := range info {
		fmt.Fprint(w, i.Id, i.Field1, i.Field2, i.Field3)
	}
}

func updateTableHandler(w http.ResponseWriter, r *http.Request) {
	model.UpdateInDB("users", "field2", "new_surname", 3)
	fmt.Fprint(w, "Information has been updated")
}

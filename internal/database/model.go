package model

// methods for interacting with the database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(tablename string, fields []string) {
	db, err := sql.Open("sqlite3", "../../internal/database/db/web_data.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()
	command := fmt.Sprintf("CREATE TABLE %s(id INTEGER PRIMARY KEY AUTOINCREMENT, %s TEXT, %s TEXT, %s TEXT);", tablename, fields[0], fields[1], fields[2])
	result, err := db.Exec(command)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

type DBData struct { // Table structure
	Id     int
	Field1 string
	Field2 string
	Field3 string
}

func GetFromDB(tablename string) []DBData {
	db, err := sql.Open("sqlite3", "../../internal/database/db/web_data.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	command := fmt.Sprintf("select * from %s", tablename)
	rows, err := db.Query(command)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer rows.Close()

	DBAllData := []DBData{}
	for rows.Next() {
		d := DBData{}
		err = rows.Scan(&d.Id, &d.Field1, &d.Field2, &d.Field3)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		DBAllData = append(DBAllData, d)
	}
	return DBAllData
}

func AddToDB(tablename string, fields []string, data []string) {
	db, err := sql.Open("sqlite3", "../../internal/database/db/web_data.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	command := fmt.Sprintf("insert into %s (%s, %s, %s) values ('%s', '%s', '%s')", tablename, fields[0], fields[1], fields[2], data[0], data[1], data[2])
	result, err := db.Exec(command)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

func UpdateInDB(tablename string, field string, text string, id int) {
	db, err := sql.Open("sqlite3", "../../internal/database/db/web_data.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	command := fmt.Sprintf("update %s set %s = '%s' where id = %d", tablename, field, text, id)
	result, err := db.Exec(command)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

func DeleteFromDB(tablename string, id int) {
	db, err := sql.Open("sqlite3", "../../internal/database/db/web_data.db")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	command := fmt.Sprintf("delete from %s where id = %d", tablename, id)
	result, err := db.Exec(command)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

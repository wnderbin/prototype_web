module web

go 1.24.0

replace web/model => ../../internal/database

require (
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	web/model v0.0.0-00010101000000-000000000000 // indirect
)

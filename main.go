package main

import (
	"test/config"
	"test/db"
	"test/http"
)

func main() {
	config.Configure()
	db.Connect()
	http.Connect()
}

package main

import (
	"techtrain_go_api/db"
	"techtrain_go_api/server"
)

func main()  {
	db.Init()

	server.Init()

	db.Close()
}
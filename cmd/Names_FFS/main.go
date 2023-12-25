package main

import (
	"ASYNC_FFS/internal/api"
	"log"
)

func main() {
	log.Println("App start")
	api.StartServer()
	log.Println("App stop")
}

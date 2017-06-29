package main

import (
	"github.com/Leafney/meiju-web/db"
	"github.com/Leafney/meiju-web/handlers"
	// "github.com/gin-gonic/gin"
	"log"
)

func init() {
	db.Connect()
}

func main() {
	router := handlers.InitRouter()
	log.Println("[info] Start Run")
	router.Run()
}

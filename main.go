package main

import (
	"github.com/Leafney/meiju-web/db"
	"github.com/Leafney/meiju-web/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	db.Connect()
}

func main() {
	router := InitRouter()
	router.Run()
}

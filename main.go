package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JammUtkarsh/cshare-server/models"
	"github.com/JammUtkarsh/cshare-server/routes"
	"github.com/JammUtkarsh/cshare-server/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	// check if database is running
	db, _ := models.CreateConnection()
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	// sets router config to prod, test, debug, etc.
	if err := utils.LoadEnv(".env"); err != nil {
		log.Fatalf("missing .env file; get sample from 'https://github.com/JammUtkarsh/cshare-server/blob/main/.env.local'")
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	routes.Routes()
	fmt.Println("Server started at port" + os.Getenv("SERVER_PORT"))
}

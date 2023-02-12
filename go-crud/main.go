package main

import (
	"context"
	"log"
	"os"

	"github.com/dblH90/go-crud/config"
	"github.com/dblH90/go-crud/driver"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var app config.GoAppTools

func main() {
	InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

	app.InfoLogger = *InfoLogger
	app.ErrorLogger = *ErrorLogger

	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Fatal("No .env file found!")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		app.ErrorLogger.Fatalln("mongodb uri string not found : ")
	}

	client := driver.Connection(uri)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			app.ErrorLogger.Fatal(err)
			return
		}
	}()

	appRouter := gin.New()
	appRouter.GET("/", func(ctx *gin.Context) {
		log.Println("Creating scalable web application with Gin!")
	})

	appRouter.GET("/users", func(ctx *gin.Context) {
		log.Println("Creating scalable web application with Gin!")
	})

	err = appRouter.Run()
	if err != nil {
		log.Fatal(err)
	}
}

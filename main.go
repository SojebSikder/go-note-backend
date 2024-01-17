package main

import (
	"github.com/SojebSikder/goframe/app/middleware"
	"github.com/SojebSikder/goframe/config"
	"github.com/SojebSikder/goframe/routes"
	orm "github.com/SojebSikder/goframe/system/core/ORM"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"context"
	"log"

	"github.com/SojebSikder/goframe/ent"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctg, _ := config.GetConfig()

	// Initialize the database connection
	hostname := ctg.Database.Hostname
	database := ctg.Database.Database
	user := ctg.Database.User
	password := ctg.Database.Password

	var databaseConnection string

	// "postgres://username:password@localhost/db_name?sslmode=disable"
	if password == "" {
		databaseConnection = "postgres://" + user + "@" + hostname + "/" + database + "?sslmode=disable"
	} else {
		databaseConnection = "postgres://" + user + ":" + password + "@" + hostname + "/" + database + "?sslmode=disable"
	}

	println(databaseConnection)
	// Initialize database connection
	client, err := ent.Open("postgres", databaseConnection)
	// client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/go-example?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//
	// Initialize ORM
	orm.Init(ctx, client)

	// Initialize the application
	r := gin.Default()

	r.Static("/static", "./"+config.StaticDir)
	r.LoadHTMLGlob(config.TemplateDir + "/*")

	// Custom middleware call here
	r.Use(middleware.Hello())
	routes.Routes(r)

	r.Run(":" + ctg.App.Port)
}

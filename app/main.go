package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	user := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	dbname := os.Getenv("APP_DB_NAME")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	// connect to postgres db
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	h := handler{
		DB: db,
	}

	router := gin.Default()
	router.GET("/dishes", h.getDishes)
	router.GET("/appetizers", h.getAppetizers)
	router.GET("/entrees", h.getEntrees)
	router.GET("/desserts", h.getDesserts)
	router.GET("/beverages", h.getBeverages)
	router.GET("/calculate", h.calculate)

	router.Run("localhost:8080")
}

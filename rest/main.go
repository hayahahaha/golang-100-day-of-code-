package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var conn *pgx.Conn

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	loadEnv()
	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	var err error

	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	router := gin.Default()

	router.GET("/", sayHello)

	router.GET("/users", listUser)

	router.Run("localhost:3000")
}

func listUser(c *gin.Context) {
	rows, _ := conn.Query(context.Background(), "select * from public.user")
	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Email, &user.password)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Some thing wrong",
			})
		}

		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func sayHello(c *gin.Context) {
	name := os.Getenv("NAME")
	fmt.Println("name in env %v", name)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	router := gin.Default()
	router.GET("/", sayHello)

	router.Run("localhost:3000")

}

func sayHello(c *gin.Context) {
	name := os.Getenv("NAME")
	fmt.Println("name in env %v", name)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}

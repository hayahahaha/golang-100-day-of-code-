package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

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

	router.GET("/users", listUser)
	router.POST("/users", addUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", removeUser)

	router.Run("localhost:3000")
}

func listUser(c *gin.Context) {
	rows, _ := conn.Query(context.Background(), "select * from users")
	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Email, &user.password); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Some thing wrong",
			})
		}

		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func addUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Some thing wrong",
		})
		return
	}

	hash, _ := HashPassword(user.password) // ignore error for the sake of simplicity
	row, err := conn.Exec(context.Background(), "insert into public.users(email, password) values($1, $2)", user.Email, hash)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Some thing wrong",
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, row)
}

func updateUser(c *gin.Context) {
	var user User
	id := c.Param("id")

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Some thing wrong",
		})
		return
	}

	row, err := conn.Exec(context.Background(), "update public.users set email=$1 where id=$2", user.Email, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Some thing wrong",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, row)
}

func removeUser(c *gin.Context) {
	id := c.Param("id")

	_, err := conn.Exec(context.Background(), "delete from public.users where $1", id)

	if err != nil {

	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Delete success",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

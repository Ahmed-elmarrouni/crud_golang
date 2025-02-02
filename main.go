package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUsers(c *gin.Context) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, email, created_at FROM public.users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	_, err := conn.Exec(context.Background(),
		"INSERT INTO public.users (name, email) VALUES ($1, $2)", newUser.Name, newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func main() {
	connectDB()

	r := gin.Default()

	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)

	go func() {
		fmt.Println(" Server running on http://localhost:8080")
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Server failed: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("\n Shutting down server...")
	CloseDB()
	fmt.Println("Server shut down successfully.")
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	ImageUrl  string    `json:"image_url"`
}

func GetUsers(c *gin.Context) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, email, created_at, image_url FROM public.users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.ImageUrl)
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
		"INSERT INTO public.users (name, email, image_url) VALUES ($1, $2, $3)", newUser.Name, newUser.Email, newUser.ImageUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updateUser User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	_, err = conn.Exec(context.Background(),
		"UPDATE public.users SET name=$1, email=$2 , image_url=$3 WHERE id=$4",
		updateUser.Name, updateUser.Email, updateUser.ImageUrl, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User di"})
		return
	}

	_, err = conn.Exec(context.Background(),
		"DELETE FROM public.users WHERE id = $1", id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to Delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}

func main() {
	connectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},        // Allow requests from your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type"},                 // Allowed headers
		AllowCredentials: true,
	}))

	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

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

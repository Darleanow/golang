// Command exercise-5b-gin implements a Todo List API using the Gin framework
// with a custom logger middleware, API-key authentication and full CRUD routes.
package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"       binding:"required"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// partialTask uses pointers so a missing JSON field stays nil and is not
// applied, distinguishing "absent" from an explicit zero value.
type partialTask struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

var (
	mu    sync.RWMutex
	tasks = map[string]Task{
		"seed-1": {ID: "seed-1", Title: "Apprendre Go", Description: "Finir les TPs", Done: false},
		"seed-2": {ID: "seed-2", Title: "Faire du sport", Done: true},
	}
)

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[%s] %s %s | %d | %s | %s",
			time.Now().Format(time.DateTime),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			c.ClientIP(),
			time.Since(start),
		)
	}
}

const apiKey = "nice-key"

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-KEY") != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "clé API invalide ou manquante"})
			return
		}
		c.Next()
	}
}

func getTasks(c *gin.Context) {
	mu.RLock()
	defer mu.RUnlock()
	list := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		list = append(list, t)
	}
	c.JSON(http.StatusOK, list)
}

func getTask(c *gin.Context) {
	id := c.Param("id")
	mu.RLock()
	defer mu.RUnlock()
	t, ok := tasks[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "tâche introuvable"})
		return
	}
	c.JSON(http.StatusOK, t)
}

func createTask(c *gin.Context) {
	var t Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = uuid.NewString()

	mu.Lock()
	tasks[t.ID] = t
	mu.Unlock()

	c.JSON(http.StatusCreated, t)
}

func updateTask(c *gin.Context) {
	id := c.Param("id")

	var patch partialTask
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	t, ok := tasks[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "tâche introuvable"})
		return
	}
	if patch.Title != nil {
		t.Title = *patch.Title
	}
	if patch.Description != nil {
		t.Description = *patch.Description
	}
	if patch.Done != nil {
		t.Done = *patch.Done
	}
	tasks[id] = t
	c.JSON(http.StatusOK, t)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	mu.Lock()
	defer mu.Unlock()
	if _, ok := tasks[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "tâche introuvable"})
		return
	}
	delete(tasks, id)
	c.Status(http.StatusNoContent)
}

func main() {
	r := gin.New()
	r.Use(loggerMiddleware())

	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTask)

	// Write routes are grouped so authMiddleware applies only to them, not to
	// the read routes above.
	protected := r.Group("/")
	protected.Use(authMiddleware())
	{
		protected.POST("/tasks", createTask)
		protected.PUT("/tasks/:id", updateTask)
		protected.DELETE("/tasks/:id", deleteTask)
	}

	log.Println("Serveur Gin démarré sur http://localhost:8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("r.Run : %v", err)
	}
}

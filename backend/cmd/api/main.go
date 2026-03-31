package main

import (
	"log"
	"taskoria_api/internal/config"
	"taskoria_api/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	var router *gin.Engine = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   "success",
			"database": "connected",
			"message":  "Taskoria API is running",
		})
	})

	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}

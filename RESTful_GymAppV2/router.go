package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func setupRouter(db *sql.DB) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// OPTIONS for CORS
	addOptionsRoute(router, "/users", "POST, PUT, DELETE, OPTIONS")
	addOptionsRoute(router, "/users/:id", "PUT, DELETE, OPTIONS")
	addOptionsRoute(router, "/coaches", "POST, PUT, DELETE, OPTIONS")
	addOptionsRoute(router, "/coaches/:id", "PUT, DELETE, OPTIONS")
	addOptionsRoute(router, "/coaches/:id/profile-pic", "POST, OPTIONS")

	// User endpoints
	router.POST("/users", createUserHandler(db))
	router.PUT("/users/:id", updateUserHandler(db))
	router.DELETE("/users/:id", deleteUserHandler(db))

	// Coach endpoints
	router.POST("/coaches", createCoachHandler(db))
	router.PUT("/coaches/:id", updateCoachHandler(db))
	router.DELETE("/coaches/:id", deleteCoachHandler(db))

	// profile upload
	router.POST("/coaches/:id/profile-pic", uploadCoachProfilePic(db))

	return router
}

func addOptionsRoute(router *gin.Engine, path string, methods string) {
	router.OPTIONS(path, func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", methods)
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Status(200)
	})
}

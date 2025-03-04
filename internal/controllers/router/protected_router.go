package router

import (
	"database/sql"
	"ice-creams-app/internal/controllers/handlers/protected-routes"
	"ice-creams-app/internal/controllers/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProtectedRoutes(r *gin.Engine, db *sql.DB, handler *protected.IceCreamHandler) {

	protected := r.Group("/api")
	protected.Use(middlewares.DBChecker(db))
	protected.Use(middlewares.JWTAuthMiddleware)
	// protected.Use(middlewares.BasicAuth)

	protected.POST("/icecreams", handler.Create)
	protected.GET("/icecreams", handler.List)
	protected.GET("/icecreams/:id", handler.Read)
	protected.PUT("/icecreams/:id", handler.Update)
	protected.DELETE("/icecreams/:id", handler.Delete)

}

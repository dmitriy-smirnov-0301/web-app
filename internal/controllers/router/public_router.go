package router

import (
	"database/sql"
	"ice-creams-app/internal/controllers/handlers/public-routes"
	"ice-creams-app/internal/controllers/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.Engine, db *sql.DB, handler *public.AuthHandler) {

	public := r.Group("/")
	public.Use(middlewares.DBChecker(db))

	public.GET("/health", handler.Health)
	public.POST("/signup", handler.Signup)
	public.POST("/login", handler.Login)
	public.POST("/update", handler.Update)
	public.POST("/recover", handler.Recover)
	public.POST("/refresh", handler.Refresh)
	public.POST("/validate", handler.Validate)

}

package router

import (
	"database/sql"
	"ice-creams-app/internal/controllers/handlers/protected-routes"
	"ice-creams-app/internal/controllers/handlers/public-routes"
	"ice-creams-app/internal/controllers/middlewares"

	"github.com/gin-gonic/gin"
)

type App struct {
	ah  *public.AuthHandler
	ih  *protected.IceCreamHandler
	adb *sql.DB
	idb *sql.DB
}

func New(ah *public.AuthHandler, ih *protected.IceCreamHandler, adb, idb *sql.DB) *App {
	return &App{
		ah:  ah,
		ih:  ih,
		adb: adb,
		idb: idb,
	}
}

func SetupRouter(app *App) *gin.Engine {

	r := gin.Default()

	r.Use(middlewares.Logger)

	RegisterPublicRoutes(r, app.adb, app.ah)
	RegisterProtectedRoutes(r, app.idb, app.ih)

	return r

}

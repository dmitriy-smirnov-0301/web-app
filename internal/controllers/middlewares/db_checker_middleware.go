package middlewares

import (
	"database/sql"
	dbchecker "ice-creams-app/internal/pkg/db-checker"

	"github.com/gin-gonic/gin"
)

func DBChecker(db *sql.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		resp := dbchecker.CheckDB(db)
		if resp.Error != nil {
			ctx.AbortWithStatusJSON(resp.StatusCode, gin.H{
				"error": resp.Error.Error(),
			})
			return
		}

		ctx.Next()

	}

}

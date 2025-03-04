package protected

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdr *IceCreamHandler) Delete(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	if err := hdr.svc.DeleteIcecreamService(id); err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ice cream deleted successfully",
	})

}

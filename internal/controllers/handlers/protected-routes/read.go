package protected

import (
	"ice-creams-app/internal/models/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdr *IceCreamHandler) Read(ctx *gin.Context) {

	id, errConv := strconv.Atoi(ctx.Param("id"))
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	icecream, err := hdr.svc.ReadIcecreamService(id)
	if err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	resp := &dto.ReadIcecreamResponse{
		Name:           icecream.Name,
		Ingredients:    icecream.Ingredients,
		ProductionDate: icecream.ProductionDate,
		BestBefore:     icecream.BestBefore,
		Price:          icecream.Price,
		Quantity:       icecream.Quantity,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})

}

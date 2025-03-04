package protected

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/models/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdr *IceCreamHandler) Update(ctx *gin.Context) {

	id, errConv := strconv.Atoi(ctx.Param("id"))
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	req := &dto.UpdateIcecreamRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	icecream := &domain.IceCream{
		ID:             id,
		Name:           req.Name,
		Ingredients:    req.Ingredients,
		ProductionDate: req.ProductionDate,
		BestBefore:     req.BestBefore,
		Price:          req.Price,
		Quantity:       req.Quantity,
	}

	if err := hdr.svc.UpdateIcecreamService(icecream); err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	resp := &dto.UpdateIcecreamResponse{
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

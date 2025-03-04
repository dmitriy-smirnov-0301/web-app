package protected

import (
	"ice-creams-app/internal/models/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (hdr *IceCreamHandler) List(ctx *gin.Context) {

	limit, errConv := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if errConv != nil || limit <= 0 {
		limit = 10
	}
	offset, errConv := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if errConv != nil || offset < 0 {
		offset = 0
	}
	sortBy := ctx.DefaultQuery("sort_by", "")
	order := strings.ToLower(ctx.DefaultQuery("order", "asc"))

	switch sortBy {
	case "production_date":
		sortBy = "production_date"
	case "best_before":
		sortBy = "best_before"
	case "price":
		sortBy = "price"
	default:
		sortBy = "id"
	}

	if order == "desc" {
		order = "DESC"
	} else {
		order = "ASC"
	}

	filter := &domain.QueryFilter{
		SortBy: sortBy,
		Order:  order,
		Limit:  limit,
		Offset: offset,
	}

	icecreams, err := hdr.svc.ListIcecreamsService(*filter)
	if err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": icecreams,
	})

}

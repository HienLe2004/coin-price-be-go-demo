package routes

import (
	_ "net/http"

	_ "github.com/HienLe2004/coin-price-be-go-demo/services/price-service/models"
	fundingrate "github.com/HienLe2004/coin-price-be-go-demo/services/price-service/services/funding_rate"
	"github.com/HienLe2004/coin-price-be-go-demo/services/price-service/services/kline"
	"github.com/gin-gonic/gin"
)

// GetFundingRate return a funding rate
// @Summary      Show a funding-rate
// @Description  get funding-rate by symbol
// @Tags         funding-rate
// @Accept       json
// @Produce      json
// @Param        symbol   	query     string  false  "symbol of funding rate"
// @Success      200		{object}  models.ResponseFundingRate
// @Router       /funding-rate [get]
func GetFundingRate(context *gin.Context) {
	fundingrate.GetFundingRate(context)
}

func GetKline(context *gin.Context) {
	kline.GetKline(context)
}

package funding_rate

import (
	"net/http"

	"github.com/HienLe2004/coin-price-be-go-demo/services/price-service/utils"
	"github.com/gin-gonic/gin"
)

func GetFundingRate(context *gin.Context) {
	var symbol = context.Query("symbol")
	if symbol == "" {
		utils.ShowError(http.StatusBadRequest, "Missing symbol", context)
		return
	}
	GetFundingRateRealTime(symbol, context)
}

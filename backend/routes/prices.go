package routes

import (
	fundingrate "github.com/HienLe2004/coin-price-be-go-demo/services/price-service/services/funding_rate"
	"github.com/HienLe2004/coin-price-be-go-demo/services/price-service/services/kline"
	"github.com/gin-gonic/gin"
)

func getFundingRate(context *gin.Context) {
	fundingrate.GetFundingRate(context)
}

func getKline(context *gin.Context) {
	kline.GetKline(context)
}

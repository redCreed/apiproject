package c_wxpay

import (
	s_wxpay "apiproject/service/wxpay"
	"apiproject/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
微信H5支付
*/
func WxH5Pay(ctx *gin.Context) {
	mwebUrl, success := s_wxpay.WxpayService.WxH5pay("标题", util.GenUniqueId().String(), 1, ctx.ClientIP())
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": map[string]interface{}{"mwebUrl": mwebUrl},
	})
}

/**
微信H5支付的回调
*/
func WxH5PayCallback(ctx *gin.Context) {
	ctx.XML(http.StatusOK, s_wxpay.WxpayService.WxH5payCallback())
}

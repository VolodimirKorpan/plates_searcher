package http

import (
	"github.com/VolodimirKorpan/go_kobi/plates"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc plates.UseCase) {
	h := NewHandler(uc)

	plates := router.Group("/plates")
	{
		plates.POST("", h.GetPlateByNumber)
	}
}
package http

import (
	"github.com/VolodimirKorpan/go_kobi/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHttpEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	autEndpoints := router.Group("/auth")
	{
		autEndpoints.POST("/sign-up", h.SignUp)
		autEndpoints.POST("/sign-in", h.SignIn)
	}
}

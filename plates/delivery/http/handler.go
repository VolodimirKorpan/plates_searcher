package http

import (
	"log"
	"net/http"

	"github.com/VolodimirKorpan/go_kobi/models"
	"github.com/VolodimirKorpan/go_kobi/plates"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase plates.UseCase
}

func NewHandler(useCase plates.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// type searchPlate struct {
// 	BeginDate string `json:"begin_date"`
// 	EndDate   string `json:"end_date"`
// 	NumPlate  string `json:"num_plate"`
// }

func (h *Handler) GetPlateByNumber(c *gin.Context) {
	inp := new(models.SearchPlate)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	plates, err := h.useCase.GetPlateByNumber(c.Request.Context(), inp.BeginDate, inp.EndDate, inp.NumPlate)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, plates)
}

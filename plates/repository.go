package plates

import (
	"context"

	"github.com/VolodimirKorpan/go_kobi/models"
)

type PlateRepository interface {
	GetPlateByNumber(ctx context.Context, plate *models.SearchPlate) ([]*models.Plate, error)
}

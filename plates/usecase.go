package plates

import (
	"context"

	"github.com/VolodimirKorpan/go_kobi/models"
)

type UseCase interface {
	GetPlateByNumber(ctx context.Context, beginDate, endDate, num string) ([]*models.Plate, error)
}

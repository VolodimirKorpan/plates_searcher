package usecase

import (
	"context"
	"fmt"

	"github.com/VolodimirKorpan/go_kobi/models"
	"github.com/VolodimirKorpan/go_kobi/plates"
	"github.com/pkg/errors"
)

type PlateUseCase struct {
	plateRepo plates.PlateRepository
}

func NewPlateUseCase(plateRepo plates.PlateRepository) *PlateUseCase {
	return &PlateUseCase{
		plateRepo: plateRepo,
	}
}

func (p PlateUseCase) GetPlateByNumber(ctx context.Context, beginDate, endDate, num string) ([]*models.Plate, error) {
	plate := &models.SearchPlate{
		BeginDate: beginDate,
		EndDate:   endDate,
		NumPlate:  num,
	}
	plateDB, err := p.plateRepo.GetPlateByNumber(ctx, plate)
	if err != nil {
		return nil, errors.Wrap(err, "p.plateRepo.GetPlateByNumber")
	}
	if plateDB == nil {
		return nil, errors.Wrap(plates.ErrPlatesNotFound, fmt.Sprintf("User '%s' not found", plate.NumPlate))
	}

	return plateDB, nil
}

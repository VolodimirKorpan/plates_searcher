package pg

import (
	"context"
	"fmt"
	"log"

	"github.com/VolodimirKorpan/go_kobi/models"
)

type PlateRepository struct {
	db *PG
}

func NewPlateRepository(db *PG) *PlateRepository {
	return &PlateRepository{
		db: db,
	}
}

func (repo *PlateRepository) GetPlateByNumber(ctx context.Context, plate *models.SearchPlate) ([]*models.Plate, error) {
	var plates []*models.Plate

	sql := fmt.Sprintf(`select id, date_time, plate, replace(pathimage, 'D:/Images/', 'http://62.122.204.137:8585/framejpg/') as image,
			camera 
			from (SELECT a.id, date_time, plate, replace(pathimage, '\', '/') as pathimage, camera 
			   FROM public.plate_events a, public.fb_camers b where plate like '%s'
				and iplates=name and date_time between '%s' and '%s') foo order by date_time `, plate.NumPlate, plate.BeginDate, plate.EndDate)

	err := repo.db.Raw(sql).Scan(&plates).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return plates, nil
}

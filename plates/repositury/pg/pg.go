package pg

import (
	"fmt"

	"github.com/VolodimirKorpan/go_kobi/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PG struct {
	*gorm.DB
}

func Dial() (*PG, error) {
	cfg := config.Get()

	if cfg.PgDB == "" {
		return nil, nil
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.PgAddr, cfg.PgUser, cfg.PgPassword, cfg.PgDB, cfg.PgPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PG{db}, nil
}
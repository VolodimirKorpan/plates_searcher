package store

import (
	"context"
	"log"

	"github.com/VolodimirKorpan/go_kobi/auth"
	"github.com/VolodimirKorpan/go_kobi/auth/repository/mysql"
	"github.com/VolodimirKorpan/go_kobi/models"
	"github.com/VolodimirKorpan/go_kobi/plates"
	"github.com/VolodimirKorpan/go_kobi/plates/repositury/pg"
	"github.com/pkg/errors"
)

type Store struct {
	MySQL *mysql.MySQL
	PG    *pg.PG

	User  auth.UserRepository
	Plate plates.PlateRepository
}

func New(ctx context.Context) (*Store, error) {
	// connect to MySQL
	log.Println("Connecting MySQL ...")
	mysqlDB, err := mysql.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "mysqldb.Dial failed")
	}
	log.Println("Connecting PostgresSQL ...")
	pgDB, err := pg.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "pgdb.Dial failed")
	}

	// run migrations
	if mysqlDB != nil {
		log.Println("Running migrations ...")
		err := mysqlDB.AutoMigrate(&models.DBUser{})
		if err != nil {
			return nil, errors.Wrap(err, "runMysqlMigrations failed")
		}
	}

	var store Store

	if mysqlDB != nil {
		store.MySQL = mysqlDB
		store.User = mysql.NewUserRepository(mysqlDB)
	}

	if pgDB != nil {
		store.PG = pgDB
		store.Plate = pg.NewPlateRepository(pgDB)
	}

	return &store, nil
}

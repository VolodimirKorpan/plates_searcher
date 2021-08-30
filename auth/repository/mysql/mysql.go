package mysql

import (
	"fmt"

	"github.com/VolodimirKorpan/go_kobi/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	*gorm.DB
}

func Dial() (*MySQL, error) {
	cfg := config.Get()
	if cfg.MysqlDB == "" {
		return nil, nil
	}
	
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlAddr, cfg.MysqlDB)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MySQL{db}, nil
}

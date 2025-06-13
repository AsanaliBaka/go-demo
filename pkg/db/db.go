package db

import (
	conifgs "go/adv-demo/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *conifgs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn))
	if err != nil {
		panic(err)
	}
	return &Db{db}
}

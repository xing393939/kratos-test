package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDBClient, NewData, NewGreeterRepo)

// Data .
type Data struct {
	db *gorm.DB
}

func NewDBClient(conf *conf.Data, logger log.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Error("mysql")
	}
	return db
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

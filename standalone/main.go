package main

import (
	"time"

	"github.com/ditrit/badaas-orm-example/standalone/models"
	"github.com/ditrit/badaas/orm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := NewGormDBConnection()
	if err != nil {
		panic(err)
	}

	err = gormDB.AutoMigrate(
		models.Product{},
		models.Company{},
		models.Seller{},
		models.Sale{},
	)
	if err != nil {
		panic(err)
	}

	crudProductService, crudProductRepository := orm.GetCRUD[models.Product, orm.UUID](gormDB)

	CreateCRUDObjects(gormDB, crudProductRepository)
	QueryCRUDObjects(crudProductService)
}

func NewGormDBConnection() (*gorm.DB, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return orm.ConnectToDialector(
		logger,
		orm.CreateDialector("localhost", "root", "postgres", "disable", "badaas_db", 26257),
		10, time.Duration(5)*time.Second,
	)
}

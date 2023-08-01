package main

import (
	"time"

	"github.com/ditrit/badaas-orm-example/fx/models"
	"github.com/ditrit/badaas/orm"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		// connect to db
		fx.Provide(NewGormDBConnection),
		// activate badaas-orm
		fx.Provide(GetModels),
		orm.AutoMigrate,

		// create crud services for models
		orm.GetCRUDServiceModule[models.Company](),
		orm.GetCRUDServiceModule[models.Product](),
		orm.GetCRUDServiceModule[models.Seller](),
		orm.GetCRUDServiceModule[models.Sale](),

		// start example data
		fx.Provide(CreateCRUDObjects),
		fx.Invoke(QueryCRUDObjects),
	).Run()
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

func GetModels() orm.GetModelsResult {
	return orm.GetModelsResult{
		Models: []any{
			models.Product{},
			models.Company{},
			models.Seller{},
			models.Sale{},
		},
	}
}

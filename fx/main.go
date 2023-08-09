package main

import (
	"github.com/ditrit/badaas-orm-example/fx/models"
	"github.com/ditrit/badaas/orm"
	"github.com/ditrit/badaas/orm/logger/gormzap"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	fx.New(
		fx.Provide(NewZapLogger),
		// connect to db
		fx.Provide(NewDBConnection),
		fx.Provide(GetModels),
		orm.AutoMigrate,

		// create crud services for models
		orm.GetCRUDServiceModule[models.Company](),
		orm.GetCRUDServiceModule[models.Product](),
		orm.GetCRUDServiceModule[models.Seller](),
		orm.GetCRUDServiceModule[models.Sale](),

		// logger for fx
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		// start example data
		fx.Provide(CreateCRUDObjects),
		fx.Invoke(QueryCRUDObjects),
	).Run()
}

func NewZapLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func NewDBConnection(zapLogger *zap.Logger) (*gorm.DB, error) {
	return orm.Open(
		postgres.Open(orm.CreateDSN("localhost", "root", "postgres", "disable", "badaas_db", 26257)),
		&gorm.Config{
			Logger: gormzap.NewDefault(zapLogger).ToLogMode(logger.Info),
		},
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

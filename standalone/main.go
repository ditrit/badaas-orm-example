package main

import (
	"github.com/ditrit/badaas-orm-example/standalone/models"
	"github.com/ditrit/badaas/orm"
	"github.com/ditrit/badaas/orm/logger"
	"github.com/ditrit/badaas/orm/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := NewDBConnection()
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

	crudProductService, crudProductRepository := orm.GetCRUD[models.Product, model.UUID](gormDB)

	CreateCRUDObjects(gormDB, crudProductRepository)
	QueryCRUDObjects(crudProductService)
}

func NewDBConnection() (*gorm.DB, error) {
	return orm.Open(
		postgres.Open(orm.CreatePostgreSQLDSN("localhost", "root", "postgres", "disable", "badaas_db", 26257)),
		&gorm.Config{
			Logger: logger.Default.ToLogMode(logger.Info),
		},
	)
}

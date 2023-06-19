package main

import (
	"fmt"
	"log"

	"github.com/ditrit/badaas-orm-example/fx/conditions"
	"github.com/ditrit/badaas-orm-example/fx/models"
	"github.com/ditrit/badaas/orm"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func CreateCRUDObjects(
	db *gorm.DB,
	crudProductRepository orm.CRUDRepository[models.Product, orm.UUID],
) ([]*models.Product, error) {
	products, err := crudProductRepository.Query(db)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		log.Println("Setting up CRUD example")

		product1 := &models.Product{
			Int: 1,
		}
		err = crudProductRepository.Create(db, product1)
		if err != nil {
			return nil, err
		}

		product2 := &models.Product{
			Int: 2,
		}
		err = crudProductRepository.Create(db, product2)
		if err != nil {
			return nil, err
		}

		company1 := &models.Company{
			Name: "ditrit",
		}
		err = db.Create(company1).Error
		if err != nil {
			return nil, err
		}
		company2 := &models.Company{
			Name: "orness",
		}
		err = db.Create(company2).Error
		if err != nil {
			return nil, err
		}

		seller1 := &models.Seller{
			Name:      "franco",
			CompanyID: &company1.ID,
		}
		err = db.Create(seller1).Error
		if err != nil {
			return nil, err
		}
		seller2 := &models.Seller{
			Name:      "agustin",
			CompanyID: &company2.ID,
		}
		err = db.Create(seller2).Error
		if err != nil {
			return nil, err
		}

		sale1 := &models.Sale{
			Product: product1,
			Seller:  seller1,
		}
		err = db.Create(sale1).Error
		if err != nil {
			return nil, err
		}
		sale2 := &models.Sale{
			Product: product2,
			Seller:  seller2,
		}
		err = db.Create(sale2).Error
		if err != nil {
			return nil, err
		}

		log.Println("Finished creating CRUD example")
		return []*models.Product{product1, product2}, nil
	}

	return nil, nil
}

func QueryCRUDObjects(
	_ []*models.Product,
	crudProductService orm.CRUDService[models.Product, orm.UUID],
	shutdowner fx.Shutdowner,
) {
	log.Println("Products with int = 1 are:")
	result, err := crudProductService.Query(
		conditions.ProductInt(orm.Eq(1)),
	)
	if err != nil {
		log.Panicln(err)
	}
	for _, product := range result {
		fmt.Printf("%+v\n", product)
	}
	shutdowner.Shutdown()
}

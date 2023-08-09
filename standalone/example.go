package main

import (
	"fmt"
	"log"

	"github.com/ditrit/badaas-orm-example/standalone/conditions"
	"github.com/ditrit/badaas-orm-example/standalone/models"
	"github.com/ditrit/badaas/orm"
	"github.com/ditrit/badaas/orm/model"
	"gorm.io/gorm"
)

func CreateCRUDObjects(
	db *gorm.DB,
	crudProductRepository orm.CRUDRepository[models.Product, model.UUID],
) ([]*models.Product, error) {
	log.Println("Migration finished, setting up CRUD example")

	return orm.Transaction(
		db,
		func(tx *gorm.DB) ([]*models.Product, error) {
			products, err := crudProductRepository.Query(tx)
			if err != nil {
				return nil, err
			}

			if len(products) == 0 {
				log.Println("Creating models")

				product1 := &models.Product{
					Int: 1,
				}
				err = crudProductRepository.Create(tx, product1)
				if err != nil {
					return nil, err
				}

				product2 := &models.Product{
					Int: 2,
				}
				err = crudProductRepository.Create(tx, product2)
				if err != nil {
					return nil, err
				}

				company1 := &models.Company{
					Name: "ditrit",
				}
				err = tx.Create(company1).Error
				if err != nil {
					return nil, err
				}
				company2 := &models.Company{
					Name: "orness",
				}
				err = tx.Create(company2).Error
				if err != nil {
					return nil, err
				}

				seller1 := &models.Seller{
					Name:      "franco",
					CompanyID: &company1.ID,
				}
				err = tx.Create(seller1).Error
				if err != nil {
					return nil, err
				}
				seller2 := &models.Seller{
					Name:      "agustin",
					CompanyID: &company2.ID,
				}
				err = tx.Create(seller2).Error
				if err != nil {
					return nil, err
				}

				sale1 := &models.Sale{
					Product: product1,
					Seller:  seller1,
				}
				err = tx.Create(sale1).Error
				if err != nil {
					return nil, err
				}
				sale2 := &models.Sale{
					Product: product2,
					Seller:  seller2,
				}
				err = tx.Create(sale2).Error
				if err != nil {
					return nil, err
				}

				log.Println("Finished creating models")
				return []*models.Product{product1, product2}, nil
			}

			log.Println("Models already created")

			return nil, nil
		},
	)
}

func QueryCRUDObjects(
	crudProductService orm.CRUDService[models.Product, model.UUID],
) {
	result, err := crudProductService.Query(
		conditions.ProductInt(orm.Eq(1)),
	)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Products with int = 1 are:")
	for _, product := range result {
		fmt.Printf("%+v\n", product)
	}
}

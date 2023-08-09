package models

import "github.com/ditrit/badaas/orm/model"

type Company struct {
	model.UUIDModel

	Name    string
	Sellers []Seller
}

type Product struct {
	model.UUIDModel

	String string
	Int    int
	Float  float64
	Bool   bool
}

type Seller struct {
	model.UUIDModel

	Name      string
	CompanyID *model.UUID
}

type Sale struct {
	model.UUIDModel

	// belongsTo Product
	Product   *Product
	ProductID model.UUID

	// belongsTo Seller
	Seller   *Seller
	SellerID model.UUID
}

package models

import (
	"github.com/ditrit/badaas/orm"
)

type Company struct {
	orm.UUIDModel

	Name    string
	Sellers []Seller
}

type Product struct {
	orm.UUIDModel

	String string
	Int    int
	Float  float64
	Bool   bool
}

type Seller struct {
	orm.UUIDModel

	Name      string
	CompanyID *orm.UUID
}

type Sale struct {
	orm.UUIDModel

	// belongsTo Product
	Product   *Product
	ProductID orm.UUID

	// belongsTo Seller
	Seller   *Seller
	SellerID orm.UUID
}

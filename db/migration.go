package main

import (
	"github.com/teacat/reiner"
)

func main() {
	db, err := reiner.New("root:@/daigou?charset=utf8")
	if err != nil {
		panic(err)
	}

	migration := db.Migration()

	migration.DropIfExists("Users")
	migration.
		Table("Users").
		Column("UserId").Char(20).Primary().
		Column("UserName").Varchar(32).
		Column("Status").Enum("pending", "qc-rejected", "qa-rejected", "unpublished", "published", "deleted").
		Create()

	migration.DropIfExists("Orders")
	migration.
		Table("Orders").
		Column("OrderID").Char(20).Primary().
		Column("Amount").Int(10).
		Column("Description").Text().
		Column("Status").Enum("pending", "qc-rejected", "qa-rejected", "unpublished", "published", "deleted").
		Create()

	migration.DropIfExists("OrdersDetail")
	migration.
		Table("OrdersDetail").
		Column("OrderID").Char(20).
		Column("Quantity").Int(10).
		Column("Amount").Int(10).
		Column("Description").Text().
		Column("ProductNumber").Char(20).
		Column("Status").Enum("pending", "qc-rejected", "qa-rejected", "unpublished", "published", "deleted").
		Create()

	migration.DropIfExists("Coupon")
	migration.
		Table("Coupon").
		Column("CouponID").Char(20).Primary().
		Column("Code").Char(10).
		Column("Name").Char(10).
		Column("Type").Enum("price", "discount").
		Column("Status").Enum("pending", "qc-rejected", "qa-rejected", "unpublished", "published", "deleted").
		Create()
}

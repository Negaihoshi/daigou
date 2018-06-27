package controllers

import (
	"fmt"

	spgateway "github.com/appleboy/go-spgateway"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/teacat/reiner"
)

type OrderController struct{}
type Order struct {
	OrderID string
	Amount  int32
}

func (u OrderController) Index(c *gin.Context) {

	db, err := reiner.New("root:@/daigou?charset=utf8")
	if err != nil {
		panic(err)
	}
	fmt.Println('t')

	var o []*Order
	db, err = db.Bind(&o).Table("Orders").Get()

	c.JSON(200, gin.H{"message": o})
	return
}

func (u OrderController) Store(c *gin.Context) {

	db, err := reiner.New("root:@/daigou?charset=utf8")
	if err != nil {
		panic(err)
	}
	guid := xid.New()

	amount := c.PostForm("Amount")
	description := c.PostForm("Description")

	store := spgateway.New(spgateway.Config{
		MerchantID: "123456",
		HashKey:    "1A3S21DAS3D1AS65D1",
		HashIV:     "1AS56D1AS24D",
	})

	order := spgateway.OrderCheckValue{
		Amt:             200,
		MerchantOrderNo: "20140901001",
		TimeStamp:       "1403243286",
		Version:         "1.2",
	}

	fmt.Println(store.OrderCheckValue(order))

	db, err = db.Table("Orders").Insert(map[string]interface{}{
		"OrderID":     guid.String(),
		"Amount":      amount,
		"Description": description,
	})

	if err != nil {
		c.JSON(500, gin.H{"message": err})
	} else {
		c.JSON(201, gin.H{"message": "order create"})
	}

	return
}

func (u OrderController) Update(c *gin.Context) {

	db, err := reiner.New("root:@/daigou?charset=utf8")
	if err != nil {
		panic(err)
	}

	orderID := c.PostForm("OrderID")
	status := c.PostForm("Status")

	db, err = db.Table("Orders").Where("OrderID", orderID).Update(map[string]interface{}{
		"Status": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"message": err})
	} else {
		c.JSON(202, gin.H{"message": orderID})
	}

	return
}

func (u OrderController) Destroy(c *gin.Context) {

	db, err := reiner.New("root:@/daigou?charset=utf8")
	if err != nil {
		panic(err)
	}

	orderID := c.PostForm("OrderID")

	db, err = db.Table("Orders").Where("OrderID", orderID).Update(map[string]interface{}{
		"Status": "deleted",
	})

	if err != nil {
		c.JSON(500, gin.H{"message": err})
	} else {
		c.JSON(202, gin.H{"message": orderID})
	}
	return
}

// func (u OrderController) Retrieve(c *gin.Context) {
// 	if c.Param("id") != "" {

// 		if err != nil {
// 			c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(200, gin.H{"message": "User founded!", "user": user})
// 		return
// 	}
// 	c.JSON(400, gin.H{"message": "bad request"})
// 	c.Abort()
// 	return
// }

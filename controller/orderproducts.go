package controller

import (
	"fmt"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Createoproduct() {
	orderpro := getinfo()
	id, err := c.Storage.OrderProducts.Insert(orderpro)
	if err != nil {
		fmt.Println("error while inserting orderproducts", err.Error())
		return
	}
	fmt.Println("id", id)
}

func (c Controller) Getlistorderproducts() {
	o, err := c.Storage.OrderProducts.GetList()
	if err != nil {
		fmt.Println("error getting list of orderproducts ", err.Error())
		return
	}

	fmt.Println(o)

}

func (c Controller) Updateorderproduct() {
	or := getinfo()
	err := c.Storage.OrderProducts.Update(or)
	if err != nil {
		fmt.Println("error updating orderproducts ")
		return
	}
	if or.Id.String() != "" {
		fmt.Println("updated")
	} else {
		fmt.Println("CREATED")
	}

}

func (c Controller) Deleteorderproducts() {
	idstr := ""
	fmt.Print("enter id to delete: ")
	fmt.Scan(&idstr)
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = c.Storage.OrderProducts.Delete(id)
	if err != nil {
		fmt.Println("error while deleting", err.Error())
		return
	}
	fmt.Println("deleted with this id:", idstr)

}

func getinfo() models.Orderproducts {
	var (
		idstr                string
		cmd, quantity, price int
		orderid, productid   string
	)
a:
	fmt.Println("enter cmd 1. CREATE 2. UPDATE ")
	fmt.Scan(&cmd)
	if cmd == 2 {
		fmt.Print("enter id to update: ")
		fmt.Scan(&idstr)
		fmt.Print("enter new price: ")
		fmt.Scan(&price)

	} else if cmd == 1 {
		fmt.Print("enter order_id: ")
		fmt.Scan(&orderid)
		fmt.Print("enter product_id: ")
		fmt.Scan(&productid)
		fmt.Print("enter quantity: ")
		fmt.Scan(&quantity)
		fmt.Print("enter price: ")
		fmt.Scan(&price)

	} else {
		fmt.Print("cmd not found ")
		goto a
	}

	if idstr != "" {
		return models.Orderproducts{
			Id:    uuid.MustParse(idstr),
			Price: price,
		}
	}

	return models.Orderproducts{
		OrderId:   uuid.MustParse(orderid),
		ProductId: uuid.MustParse(productid),
		Quantity:  quantity,
		Price:     price,
	}
}

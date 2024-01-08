package controller

import (
	"fmt"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Createproducts() {
	product := getproductinfo()
	id, err := c.Storage.ProductsStorage.Insert(product)
	if err != nil {
		fmt.Println("error inserting category ")
		return
	}
	fmt.Println("id ", id)

}

func (c Controller) Getprpoducts() {
	products, err := c.Storage.ProductsStorage.Getlist()
	if err != nil {
		fmt.Println("error list products ", err.Error())
		return
	}
	fmt.Println(products)
}
func (c Controller) Updateproducts() {
	update := getproductinfo()
	err := c.Storage.ProductsStorage.Update(update)
	if err != nil {
		fmt.Println("error updating products", err.Error())
		return
	}
	if update.Id.String() != "" {
		fmt.Println("UPDATED")
	} else {
		fmt.Println("CREATED")
	}

}

func (c Controller) Deleteproducts() {
	idstr := ""
	fmt.Print("enter id: ")
	fmt.Scan(&idstr)
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while pasing in 50 ", err.Error())
		return
	}
	err = c.Storage.ProductsStorage.Delete(id)
	fmt.Println("delete7d products with this id: ", idstr)

}

func getproductinfo() models.Products {
	var (
		idstr, name string
		price, cmd  int
	)
a:
	fmt.Println("enter command : 1.create 2. update ")
	fmt.Scan(&cmd)
	if cmd == 2 {
		fmt.Print("enter id to update : ")
		fmt.Scan(&idstr)
		fmt.Print("enter new price: ")
		fmt.Scan(&price)
	} else if cmd == 1 {
		fmt.Print("enter name : ")
		fmt.Scan(&name)
		fmt.Print("enter price: ")
		fmt.Scan(&price)
	} else {
		fmt.Println("not found cmd ")
		goto a
	}
	if idstr != "" {
		return models.Products{
			Id:    uuid.MustParse(idstr),
			Name:  name,
			Price: price,
		}
	}
	return models.Products{
		Name:  name,
		Price: price,
	}

}

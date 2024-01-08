package controller

import (
	"fmt"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Createorders() {
	order := getorderinfo()
	id, err := c.Storage.OrdersStorage.Insert(order)
	if err != nil {
		fmt.Println("error inserting orders ", err.Error())
		return
	}
	fmt.Println("created with this id: ", id)
}

func (c Controller) GetListorders() {
	orders, err := c.Storage.OrdersStorage.GetList()
	if err != nil {
		fmt.Println("error getting list orders ", err.Error())
		return
	}
	fmt.Println(orders)

}

func (c Controller) Updateorders() {
	order := getorderinfo()
	err := c.Storage.OrdersStorage.Update(order)
	if err != nil {
		fmt.Println("error ", err.Error())
		return
	}

}

func (c Controller) Deleteorders() {
	idstr := ""
	fmt.Println("enter id to enter :")
	fmt.Scan(&idstr)
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while parsing in line 46", err.Error())
		return
	}
	err = c.Storage.OrdersStorage.Delete(id)
	if err != nil {
		fmt.Println("error while deleting orders ", err.Error())
		return
	}
	fmt.Println("deleted orders with this id ", idstr)

}
func getorderinfo() models.Orders {
	var (
		idstr       string
		amount, cmd int
		userIdstr   string
	)
a:
	fmt.Println("enter cmd: 1.create 2.update")
	fmt.Scan(&cmd)
	if cmd == 2 {
		fmt.Print("enter id to update ")
		fmt.Scan(&idstr)
		fmt.Print("enter new amount: ")
		fmt.Scan(&amount)
	} else if cmd == 1 {
		fmt.Print("enter amount: ")
		fmt.Scan(&amount)
		fmt.Print("enter user id:")
		fmt.Scan(&userIdstr)
	} else {
		fmt.Print("cmd not found ")
		goto a
	}
	if idstr != "" {
		return models.Orders{
			Id:     uuid.MustParse(idstr),
			Amount: amount,
		}
	}
	return models.Orders{
		Amount: amount,
		UserId: uuid.MustParse(userIdstr),
	}

}

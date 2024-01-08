package controller

import (
	"fmt"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) CreateUsers() {
	user := getUserInfo()
	id, err := c.Storage.UsersStorage.Insert(user)
	if err != nil {
		fmt.Println("error while inserting ", err.Error())
		return
	}
	fmt.Println("id created : ", id)

}

func (c Controller) GetListUsers() {
	user, err := c.Storage.UsersStorage.GetList()
	if err != nil {
		fmt.Println("error list user ", err.Error())
		return
	}
	fmt.Println(user)

}

func (c Controller) UpdateUser() {
	user := getUserInfo()
	err := c.Storage.UsersStorage.Update(user)
	if err != nil {
		fmt.Println("error updating user ", err.Error())
		return
	}
	if user.Id.String() != "" {
		fmt.Println("UPDATED")
	} else {
		fmt.Println("CREATED")
	}

}

func (c Controller) Deleteusers() {
	idstr := ""

	fmt.Print("enter id ")
	fmt.Scan(&idstr)
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while parsing 50 ", err.Error())
		return
	}
	err = c.Storage.UsersStorage.Delete(id)
	if err != nil {
		fmt.Println("error while deleting user", err.Error())
		return
	}
	fmt.Println("deleted user with this id :", idstr)

}

func getUserInfo() models.Users {
	var (
		idstr, email, phone string
		name                string
		cmd                 int
	)
a:
	fmt.Println("enter cmd: 1. CREATE 2. UPDATE ")
	fmt.Scan(&cmd)
	if cmd == 2 {
		fmt.Print("enter id to update : ")
		fmt.Scan(&idstr)
		fmt.Print("enter new name: ")
		fmt.Scan(&name)
	} else if cmd == 1 {
		fmt.Print("enter name : ")
		fmt.Scan(&name)
		fmt.Print("enter email: ")
		fmt.Scan(&email)
		fmt.Print("enter phone: ")
		fmt.Scan(&phone)
	} else {
		fmt.Println("cmd not found ")
		goto a
	}
	if idstr != "" {
		return models.Users{
			Id:        uuid.MustParse(idstr),
			Firstname: name,
		}
	}
	return models.Users{
		Firstname: name,
		Email:     email,
		Phone:     phone,
	}
}

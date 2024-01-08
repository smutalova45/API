package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
	"main.go/pkg/check"
)

/*
post-create*
put-update
get-getbyid*
delete-delete
*/
func (c Controller) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateUsers(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdUser(w, r)
		} else {
			c.GetListUsers(w, r)
		}
	case http.MethodPut:
		c.UpdateUser(w, r)
	case http.MethodDelete:
		c.Deleteusers(w, r)
	}

}

func (c Controller) CreateUsers(w http.ResponseWriter, r *http.Request) {
	//strukturaga saqlash
	user := models.Users{}
	//kegan malumotni oqish structga
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("error while reading data to users ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}

	if !check.PhoneNumber(user.Phone) {
		fmt.Println("the phone number format is not correct!")
		hanldeResponse(w, http.StatusBadRequest, errors.New("phone type is not correct!"))
		return
	}

	id, err := c.Storage.UsersStorage.Insert(user)
	if err != nil {

		fmt.Println("error while inserting user", err.Error())
		hanldeResponse(w, 500, err.Error())

		return
	}

	id1, err := uuid.Parse(id)

	resp, err := c.Storage.UsersStorage.GetById(id1)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())
	}

	hanldeResponse(w, http.StatusCreated, resp)

}

func (c Controller) GetListUsers(w http.ResponseWriter, r *http.Request) {

	user, err := c.Storage.UsersStorage.GetList()
	if err != nil {
		fmt.Println("error whilegetting list of users ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, http.StatusOK, user)

}

func (c Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.Users{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("error while reading user to update ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	err := c.Storage.UsersStorage.Update(user)
	if err != nil {
		fmt.Println("error updating user ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}

	hanldeResponse(w, 200, user.Id)

}

func (c Controller) Deleteusers(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values["id"][0]
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while parsing user id to delete ", err.Error())
		hanldeResponse(w,500,err.Error())
		return
	}

	err = c.Storage.UsersStorage.Delete(id)

	if err != nil {
		fmt.Println("error while deleting user", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, idstr)

}

func (c Controller) GetByIdUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	fmt.Println(values)
	idstr := values["id"][0]

	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w,500,err)
		return
	}
	user, err := c.Storage.UsersStorage.GetById(id)
	if err != nil {
		fmt.Println("error while getting id user", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	js, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error while marsheling user", err.Error())
		hanldeResponse(w,500,err.Error())
		return
	}
	hanldeResponse(w, 200, js)

}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) OrderProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.Createorderproduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdOrderproducts(w, r)
		} else {
			c.Getlistorderproducts(w, r)
		}
	case http.MethodPut:
		c.Updateorderproduct(w, r)
	case http.MethodDelete:
		c.Deleteorderproducts(w, r)
	}

}
func (c Controller) Createorderproduct(w http.ResponseWriter, r *http.Request) {
	orp := models.Orderproducts{}
	if err := json.NewDecoder(r.Body).Decode(&orp); err != nil {
		fmt.Println("error while reading data to orderproducts", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	id, err := c.Storage.OrderProductsStorage.Insert(orp)
	if err != nil {
		fmt.Println("error while inserting orderproducts", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	idd, err := uuid.Parse(id)

	resp, err := c.Storage.OrderProductsStorage.GetById(idd)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, resp)
}

func (c Controller) Getlistorderproducts(w http.ResponseWriter, R *http.Request) {

	o, err := c.Storage.OrderProductsStorage.GetList()
	if err != nil {
		fmt.Println("error getting list of orderproducts ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}

	hanldeResponse(w, http.StatusOK, o)

}

func (c Controller) Updateorderproduct(w http.ResponseWriter, r *http.Request) {
	or := models.Orderproducts{}
	if err := json.NewDecoder(r.Body).Decode(&or); err != nil {
		fmt.Println("error while reading data to update orderproducts", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	err := c.Storage.OrderProductsStorage.Update(or)
	if err != nil {
		fmt.Println("error updating orderproducts ")
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, or.Id)

}

func (c Controller) Deleteorderproducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values["id"][0]
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err)
		return
	}
	err = c.Storage.OrderProductsStorage.Delete(id)
	if err != nil {
		fmt.Println("error while deleting orderproducts", err.Error())
		hanldeResponse(w, 500, err)
		return
	}
	hanldeResponse(w, 200, idstr)

}
func (c Controller) GetByIdOrderproducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values["id"][0]
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err)
		return
	}

	op, err := c.Storage.OrderProductsStorage.GetById(id)
	if err != nil {
		fmt.Println("error while getting by id orderproducts", err.Error())
		hanldeResponse(w, 500, err)
		return
	}
	js, err := json.Marshal(op)
	if err != nil {
		fmt.Println("error while marsheling data to orderproducts", err.Error())
		hanldeResponse(w, 500, err)
		return
	}
	hanldeResponse(w, 200, js)

}

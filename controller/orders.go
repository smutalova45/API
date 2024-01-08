package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Orders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.Createorders(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdOrders(w, r)
		} else {
			c.GetListorders(w, r)
		}
	case http.MethodPut:
		c.Updateorders(w, r)
	case http.MethodDelete:
		c.Deleteorders(w, r)
	}

}
func (c Controller) Createorders(w http.ResponseWriter, r *http.Request) {
	orders := models.Orders{}
	if err := json.NewDecoder(r.Body).Decode(&orders); err != nil {
		fmt.Println("error while reading data to orders", err.Error())
		hanldeResponse(w, 500, err.Error())
	}
	id, err := c.Storage.OrdersStorage.Insert(orders)
	if err != nil {
		fmt.Println("error inserting orders ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	id1, err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp, err := c.Storage.OrdersStorage.GetById(id1)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())

	}
	hanldeResponse(w, http.StatusCreated, resp)
}

func (c Controller) GetListorders(w http.ResponseWriter, r *http.Request) {

	orders, err := c.Storage.OrdersStorage.GetList()
	if err != nil {
		fmt.Println("error while getting list orders ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, http.StatusCreated, orders)

}

func (c Controller) Updateorders(w http.ResponseWriter, r *http.Request) {
	order := models.Orders{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println("error while updating orders", err.Error())
		hanldeResponse(w, 500, err.Error())
	}
	err := c.Storage.OrdersStorage.Update(order)
	if err != nil {
		fmt.Println("error ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, order.Id)

}

func (c Controller) Deleteorders(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values["id"][0]
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while parsing in line order id", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	err = c.Storage.OrdersStorage.Delete(id)
	if err != nil {
		fmt.Println("error while deleting orders ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, idstr)

}

func (c Controller) GetByIdOrders(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values["id"][0]
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}

	o, err := c.Storage.OrdersStorage.GetById(id)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	js, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, js)

}

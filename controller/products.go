package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"main.go/models"
)

func (c Controller) Products(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.Createproducts(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetByIdProducts(w, r)
		} else {
			c.GetListproducts(w, r)
		}
	case http.MethodPut:
		c.Updateproducts(w, r)
	case http.MethodDelete:
		c.Deleteproducts(w, r)
	}
}
func (c Controller) Createproducts(w http.ResponseWriter, r *http.Request) {
	product := models.Products{}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println("error while reading data to products ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	id, err := c.Storage.ProductsStorage.Insert(product)
	if err != nil {
		fmt.Println("error inserting product", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	id1, err := uuid.Parse(id)

	resp, err := c.Storage.ProductsStorage.GetById(id1)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, resp)

}

func (c Controller) GetListproducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.Storage.ProductsStorage.Getlist()
	if err != nil {
		fmt.Println("error while getting list products ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, products)
}
func (c Controller) Updateproducts(w http.ResponseWriter, r *http.Request) {
	product := models.Products{}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println("error while updting products", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	err := c.Storage.ProductsStorage.Update(product)
	if err != nil {
		fmt.Println("error while updating products", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, http.StatusAccepted, product.Id)

}

func (c Controller) Deleteproducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values["id"][0]
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("error while parsing id ", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	err = c.Storage.ProductsStorage.Delete(id)
	if err != nil {
		fmt.Println("error while deleting products", err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	hanldeResponse(w, 200, idstr)

}

func (c Controller) GetByIdProducts(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	idstr := values["id"][0]
	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	p, err := c.Storage.ProductsStorage.GetById(id)
	if err != nil {
		fmt.Println(err.Error())
		hanldeResponse(w, 500, err.Error())
		return
	}
	js, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error while marshelling", err.Error())
		hanldeResponse(w, 500, err.Error())
	}
	hanldeResponse(w, 200, js)

}

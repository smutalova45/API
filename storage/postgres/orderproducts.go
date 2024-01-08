package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"main.go/models"
)

type orderproductsRepo struct {
	DB *sql.DB
}

func NewOrderproductsRepo(db *sql.DB) orderproductsRepo {
	return orderproductsRepo{
		DB: db,
	}
}

func (or orderproductsRepo) Insert(ord models.Orderproducts) (string, error) {
	id := uuid.New()
	if _, err := or.DB.Exec(`insert into orderproducts values($1, $2, $3, $4, $5)`, id, ord.OrderId, ord.ProductId, ord.Quantity, ord.Price); err != nil {
		return "", err
	}
	return id.String(), nil

}

func (or orderproductsRepo) GetList() ([]models.Orderproducts, error) {
	rows, err := or.DB.Query(`select * from orderproducts `)
	if err != nil {
		return nil, err
	}
	ord1 := []models.Orderproducts{}
	for rows.Next() {
		ord := models.Orderproducts{}
		if err = rows.Scan(&ord.Id, &ord.OrderId, &ord.ProductId, &ord.Quantity, &ord.Price); err != nil {
			return nil, err
		}
		ord1 = append(ord1, ord)

	}
	return ord1, nil

}

func (or orderproductsRepo) Update(ord models.Orderproducts) error {
	_, err := or.DB.Exec(`update orderproducts set price= $1  where id = $2`, ord.Price, ord.Id)
	if err != nil {
		return err
	}
	return nil

}

func (or orderproductsRepo) Delete(id uuid.UUID) error {
	if _, err := or.DB.Exec(`delete from orderproducts where id =$1`, id); err != nil {
		return err
	}
	return nil

}
func (or orderproductsRepo) GetById(id uuid.UUID) (models.Orderproducts, error) {
	orderproduct := models.Orderproducts{}
	if err := or.DB.QueryRow(`select order id , order_id , product_id, quantity, price from orderproducts where id=$1`, id).Scan(
		&orderproduct.Id,
		&orderproduct.OrderId,
		&orderproduct.ProductId,
		&orderproduct.Quantity,
		&orderproduct.Price,
	); err != nil {
		return models.Orderproducts{}, err
	}
	return orderproduct, nil

}

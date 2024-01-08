package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"main.go/models"
)

type ordersRepo struct {
	DB *sql.DB
}

func NewOrdersRepo(db *sql.DB) ordersRepo {
	return ordersRepo{
		DB: db,
	}
}

func (o ordersRepo) Insert(orders models.Orders) (string, error) {
	id := uuid.New()
	orders.CreatedAt = time.Now()
	if _, err := o.DB.Exec(`insert into orders values($1, $2, $3, $4 )`, id, orders.Amount, orders.CreatedAt, orders.UserId); err != nil {
		return "", err
	}
	return id.String(), nil

}

func (o ordersRepo) GetList() ([]models.Orders, error) {
	rows, err := o.DB.Query(`select * from orders `)
	if err != nil {
		return nil, err
	}
	o1 := []models.Orders{}
	for rows.Next() {
		or := models.Orders{}
		if err = rows.Scan(&or.Id, &or.Amount, &or.CreatedAt, &or.UserId); err != nil {
			return nil, err
		}
		o1 = append(o1, or)

	}
	return o1, nil

}

func (o ordersRepo) Update(ord models.Orders) error {
	_, err := o.DB.Exec(`update orders set amount=$1 where id=$2`, ord.Amount, ord.Id)
	if err != nil {
		return err
	}
	return nil

}

func (o ordersRepo) Delete(id uuid.UUID) error {

	if _, err := o.DB.Exec("delete from orderproducts where order_id = $1", id); err != nil {
		return err
	}

	if _, err := o.DB.Exec(`delete from orders where id = $1 `, id); err != nil {
		return err
	}
	return nil
}

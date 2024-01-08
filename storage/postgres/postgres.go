package postgres

import (
	"database/sql"
	"fmt"

	"main.go/config"
)

type Storage struct {
	DB              *sql.DB
	UsersStorage    usersRepo
	OrdersStorage   ordersRepo
	ProductsStorage productsRepo
	OrderProductsStorage   orderproductsRepo
}

func New(cfg config.Config) (Storage, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Storage{}, err
	}
	usersRepo := NewUsersRepo(db)
	ordersRepo := NewOrdersRepo(db)
	productsRepo := NewProductsRepo(db)
	orderproductsRepo := NewOrderproductsRepo(db)
	return Storage{
		DB:              db,
		UsersStorage:    usersRepo,
		ProductsStorage: productsRepo,
		OrdersStorage:   ordersRepo,
		OrderProductsStorage:   orderproductsRepo,
	}, nil

}

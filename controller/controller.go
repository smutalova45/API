package controller

import "main.go/storage/postgres"



type Controller struct {
	Storage postgres.Storage
}

func New(storage postgres.Storage) Controller {
	return Controller{
		Storage: storage,
	}
}

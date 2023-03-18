package models

import "github.com/docker/distribution/uuid"

type Product struct {
	Id          string
	Name        string
	Description string
	Price       float64
}

func CreateProduct(name string, description string, price float64) *Product {
	return &Product{
		Id: uuid.New().String,
		Name: name,
		Price: price
	}
}
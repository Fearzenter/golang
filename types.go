package main

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type super struct {
	Id        int
	Name      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type ProductRepository interface {
	CreateProduct(p Product) (Product, error)
	GetProductByID(id int) (Product, error)
	UpdateProductById(id int, p Product) (Product, error)
	DeleteProductByID(id int) (Product, error)
}

type JsonStorage struct {
	name     string
	fileName string
}

func (j JsonStorage) marshalFile(target interface{}) error {
	encode, err := json.Marshal(target)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(j.fileName, encode, 0644)
}

func (j JsonStorage) UnmarshalFile(target interface{}) error {
	decode, err := ioutil.ReadFile(j.fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(decode, target)
}

func (j JsonStorage) CreateProduct(p Product) (Product, error) {
	var products []Product
	if err := j.UnmarshalFile(&products); err != nil {
		return Product{}, err
	}
	products = append(products, p)
	if err := j.marshalFile(products); err != nil {
		return Product{}, errors.New("marshal fail")
	}
	fmt.Println(p.Name, "is created")
	return p, nil
}

func (j JsonStorage) GetProductByID(id int) (Product, error) {
	var products []Product
	if err := j.UnmarshalFile(&products); err != nil {
		return Product{}, errors.New("unmarshal fail")
	}
	fmt.Println(products[id].Name, "is:")
	var index int
	for i, p := range products {
		if p.Id == id {
			index = i
		}
	}
	return products[index], nil
}

func (j JsonStorage) UpdateProductById(id int, p Product) ([]Product, error) {
	var products []Product
	if err := j.UnmarshalFile(&products); err != nil {
		return []Product{}, err
	}
	if err, _ := j.DeleteProductByID(id); err != nil {
		return []Product{}, errors.New("delete fail")
	}
	products = append(products, p)
	return products, nil
}

func (j JsonStorage) DeleteProductByID(id int) ([]Product, error) { //нужно вернуть массив новых продуктов или уд. продукт?
	var products []Product

	var index int
	for i, p := range products {
		if p.Id == id {
			index = i
		}
	}

	if err := j.UnmarshalFile(&products); err != nil {
		return []Product{}, err
	}
	if index >= 0 && index < len(products) {
		products = append(products[:index], products[index+1:]...)
	}
	if err := j.marshalFile(products); err != nil {
		return []Product{}, errors.New("delete fail")
	}
	fmt.Println("Product deleted")
	return products, nil
}

type Store struct {
	super
	ImageLink string
}
type StoreType struct {
	super
	ImageLink string
}
type Product struct {
	super
	ImageLinks list.List
}
type ProductType struct {
	super
	ImageLink string
	StoreId   int
}
type ProductTypeProduct struct {
	ProductTypeId int
	ProductId     int
}
type StoreTypeStore struct {
	StoreTypeId int
	StoreId     int
}

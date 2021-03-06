package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type super struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ProductRepository interface {
	CreateProduct(p Product) (Product, error)
	GetProductByID(id int) (Product, error)
	UpdateProductById(id int, p Product) (Product, error)
	DeleteProductByID(id int) (Product, error)
}

type JsonStorage struct {
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
	/*err1 := json.Unmarshal(decode, target)
	if _, ok := err1.(*json.MarshalerError); ok {
		return nil
	}*/
	if len(decode) == 0 {
		return nil
	}
	return json.Unmarshal(decode, target)
}

//Product crud
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

	var index int
	for i, p := range products {
		if p.Id == id {
			index = i
		}
	}

	products[index].Name = p.Name
	products[index].UpdatedAt = time.Now()
	products[index].ImageLinks = p.ImageLinks

	products = append(products, p)
	return products, nil
}

func (j JsonStorage) DeleteProductByID(id int) (Product, error) { //???????????????????? ?????????????????? ??????????????
	var products []Product

	if err := j.UnmarshalFile(&products); err != nil {
		return Product{}, err
	}

	var index int //?????????? ?????????????? ???? ????????
	for i, p := range products {
		if p.Id == id {
			index = i
		}
	}
	deleted := products[index]               //???????????????????? ???? ????????????????
	if index >= 0 && index < len(products) { //????????????????
		products = append(products[:index], products[index+1:]...)
	}
	if err := j.marshalFile(products); err != nil { //????????????????????
		return Product{}, errors.New("delete fail")
	}
	fmt.Println("Product deleted")
	return deleted, nil
}

//Store crud
func (j JsonStorage) CreateStore(s Store) (Store, error) {
	var stores []Store
	if err := j.UnmarshalFile(&stores); err != nil {
		return Store{}, err
	}
	stores = append(stores, s)
	if err := j.marshalFile(stores); err != nil {
		return Store{}, errors.New("marshal fail")
	}
	fmt.Println(s.Name, "is created")
	return s, nil
}

func (j JsonStorage) GetStoreByID(id int) (Store, error) {
	var stores []Store
	if err := j.UnmarshalFile(&stores); err != nil {
		return Store{}, errors.New("unmarshal fail")
	}
	fmt.Println(stores[id].Name, "is:")
	var index int
	for i, s := range stores {
		if s.Id == id {
			index = i
		}
	}
	return stores[index], nil
}

func (j JsonStorage) UpdateStoreById(id int, s Store) ([]Store, error) {
	var stores []Store
	if err := j.UnmarshalFile(&stores); err != nil {
		return []Store{}, err
	}

	var index int
	for i, s := range stores {
		if s.Id == id {
			index = i
		}
	}

	stores[index].Name = s.Name
	stores[index].UpdatedAt = time.Now()
	stores[index].ImageLink = s.ImageLink
	stores = append(stores, s)
	return stores, nil
}

func (j JsonStorage) DeleteStoreByID(id int) ([]Store, error) { //?????????? ?????????????? ???????????? ?????????? ?????????????????? ?????? ????. ???????????????
	var stores []Store

	var index int
	for i, s := range stores {
		if s.Id == id {
			index = i
		}
	}

	if err := j.UnmarshalFile(&stores); err != nil {
		return []Store{}, err
	}
	if index >= 0 && index < len(stores) {
		stores = append(stores[:index], stores[index+1:]...)
	}
	if err := j.marshalFile(stores); err != nil {
		return []Store{}, errors.New("delete fail")
	}
	fmt.Println("Product deleted")
	return stores, nil
}

//StoreType crud
func (j JsonStorage) CreateStoreType(s StoreType) (StoreType, error) {
	var storeTypes []StoreType
	if err := j.UnmarshalFile(&storeTypes); err != nil {
		return StoreType{}, err
	}
	storeTypes = append(storeTypes, s)
	if err := j.marshalFile(storeTypes); err != nil {
		return StoreType{}, errors.New("marshal fail")
	}
	fmt.Println(s.Name, "is created")
	return s, nil
}

func (j JsonStorage) GetStoreTypeByID(id int) (StoreType, error) {
	var storeTypes []StoreType
	if err := j.UnmarshalFile(&storeTypes); err != nil {
		return StoreType{}, errors.New("unmarshal fail")
	}
	fmt.Println(storeTypes[id].Name, "is:")
	var index int
	for i, s := range storeTypes {
		if s.Id == id {
			index = i
		}
	}
	return storeTypes[index], nil
}

func (j JsonStorage) UpdateStoreTypeById(id int, s StoreType) ([]StoreType, error) {
	var storeTypes []StoreType
	if err := j.UnmarshalFile(&storeTypes); err != nil {
		return []StoreType{}, err
	}

	var index int
	for i, s := range storeTypes {
		if s.Id == id {
			index = i
		}
	}

	storeTypes[index].Name = s.Name
	storeTypes[index].UpdatedAt = time.Now()
	storeTypes[index].ImageLink = s.ImageLink

	storeTypes = append(storeTypes, s)
	return storeTypes, nil
}

func (j JsonStorage) DeleteStoreTypeByID(id int) ([]StoreType, error) { //?????????? ?????????????? ???????????? ?????????? ?????????????????? ?????? ????. ???????????????
	var storeTypes []StoreType

	var index int
	for i, s := range storeTypes {
		if s.Id == id {
			index = i
		}
	}

	if err := j.UnmarshalFile(&storeTypes); err != nil {
		return []StoreType{}, err
	}
	if index >= 0 && index < len(storeTypes) {
		storeTypes = append(storeTypes[:index], storeTypes[index+1:]...)
	}
	if err := j.marshalFile(storeTypes); err != nil {
		return []StoreType{}, errors.New("delete fail")
	}
	fmt.Println("Product deleted")
	return storeTypes, nil
}

//ProductType crud
func (j JsonStorage) CreateProductType(p ProductType) (ProductType, error) {
	var productTypes []ProductType
	if err := j.UnmarshalFile(&productTypes); err != nil {
		return ProductType{}, err
	}
	productTypes = append(productTypes, p)
	if err := j.marshalFile(productTypes); err != nil {
		return ProductType{}, errors.New("marshal fail")
	}
	fmt.Println(p.Name, "is created")
	return p, nil
}

func (j JsonStorage) GetProductTypeByID(id int) (ProductType, error) {
	var productTypes []ProductType
	if err := j.UnmarshalFile(&productTypes); err != nil {
		return ProductType{}, errors.New("unmarshal fail")
	}
	fmt.Println(productTypes[id].Name, "is:")
	var index int
	for i, s := range productTypes {
		if s.Id == id {
			index = i
		}
	}
	return productTypes[index], nil
}

func (j JsonStorage) UpdateProductTypeById(id int, p ProductType) ([]ProductType, error) {
	var productTypes []ProductType
	if err := j.UnmarshalFile(&productTypes); err != nil {
		return []ProductType{}, err
	}

	var index int
	for i, p := range productTypes {
		if p.Id == id {
			index = i
		}
	}

	productTypes[index].Name = p.Name
	productTypes[index].UpdatedAt = time.Now()
	productTypes[index].ImageLink = p.ImageLink

	productTypes = append(productTypes, p)
	return productTypes, nil
}

func (j JsonStorage) DeleteProductTypeByID(id int) ([]ProductType, error) { //?????????? ?????????????? ???????????? ?????????? ?????????????????? ?????? ????. ???????????????
	var productTypes []ProductType

	var index int
	for i, s := range productTypes {
		if s.Id == id {
			index = i
		}
	}

	if err := j.UnmarshalFile(&productTypes); err != nil {
		return []ProductType{}, err
	}
	if index >= 0 && index < len(productTypes) {
		productTypes = append(productTypes[:index], productTypes[index+1:]...)
	}
	if err := j.marshalFile(productTypes); err != nil {
		return []ProductType{}, errors.New("delete fail")
	}
	fmt.Println("Product deleted")
	return productTypes, nil
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
	ImageLinks []string
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

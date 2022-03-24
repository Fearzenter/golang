package main

type super struct {
	id        int
	name      string
	createdAt string
	updatedAt string
	deletedAt string
}

type store struct {
	super
	imageLink string
}
type storeType struct {
	super
	imageLink string
}
type product struct {
	super
	imageLinks [4]string
}
type productType struct {
	super
	imageLink string
	storeId   int
}
type productTypeProduct struct {
	productTypeId int
	productId     int
}
type storeTypeStore struct {
	storeTypeId int
	storeId     int
}

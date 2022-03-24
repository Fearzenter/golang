package main

type super struct {
	Id        int
	Name      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type store struct {
	super
	ImageLink string
}
type storeType struct {
	super
	ImageLink string
}
type product struct {
	super
	ImageLinks [4]string
}
type productType struct {
	super
	ImageLink string
	StoreId   int
}
type productTypeProduct struct {
	ProductTypeId int
	ProductId     int
}
type storeTypeStore struct {
	StoreTypeId int
	StoreId     int
}

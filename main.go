package main

import (
	"fmt"
)

func test() {
	{
		testStore := store{super: super{id: -1, name: "Vyasoliy_molochnik", createdAt: "22.03.2022", updatedAt: "22.03.2022", deletedAt: "22.03.2022"}, imageLink: "yandex.ru"}
		fmt.Println(testStore)
	} //test store
	{
		testStoreType := storeType{super: super{id: -1, name: "productoviy", createdAt: "22.03.2022", updatedAt: "22.03.2022", deletedAt: "22.03.2022"}, imageLink: "yandex.ru"}
		fmt.Println(testStoreType)
	} //test storeType
	{
		links := [...]string{
			"yandex.ru",
			"google.com",
			"vk.com",
			"youtube.com",
		}
		testProduct := product{super: super{id: -1, name: "milk", createdAt: "22.03.2022", updatedAt: "22.03.2022", deletedAt: "22.03.2022"}, imageLinks: links}
		fmt.Println(testProduct)
	} //test product
	{
		testProductType := productType{super: super{id: -1, name: "molochka", createdAt: "22.03.2022", updatedAt: "22.03.2022", deletedAt: "22.03.2022"}, imageLink: "yandex.ru"}
		fmt.Println(testProductType)
	} //test productType
	{
		testProductTypeProduct := productTypeProduct{productTypeId: -1, productId: -1}
		fmt.Println(testProductTypeProduct)
	} //test productTypeProduct
	{
		testStoreTypeStore := storeTypeStore{storeTypeId: -1, storeId: -1}
		fmt.Println(testStoreTypeStore)
	} //test storeTypeStore
}

func main() {
	fmt.Println("Hello Golang!")
	test()
}

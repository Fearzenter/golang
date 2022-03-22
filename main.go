package main

import (
	"fmt"
)

type super struct {
	id         int
	name       string
	created_at string
	updated_at string
	deleted_at string
}

type store struct {
	super      super
	image_link string
}
type store_type struct {
	super      super
	image_link string
}
type product struct {
	super       super
	image_links [4]string
}
type product_type struct {
	super      super
	image_link string
	store_id   int
}
type product_type_product struct {
	product_type_id int
	product_id      int
}
type store_type_store struct {
	store_type_id int
	store_id      int
}

func test() {
	{
		test_store := store{super: super{id: -1, name: "Vyasoliy_molochnik", created_at: "22.03.2022", updated_at: "22.03.2022", deleted_at: "22.03.2022"}, image_link: "yandex.ru"}
		fmt.Println(test_store)
	} //test store
	{
		test_store_type := store_type{super: super{id: -1, name: "productoviy", created_at: "22.03.2022", updated_at: "22.03.2022", deleted_at: "22.03.2022"}, image_link: "yandex.ru"}
		fmt.Println(test_store_type)
	} //test store_type
	{
		links := [...]string{
			"yandex.ru",
			"google.com",
			"vk.com",
			"youtube.com",
		}
		test_product := product{super: super{id: -1, name: "milk", created_at: "22.03.2022", updated_at: "22.03.2022", deleted_at: "22.03.2022"}, image_links: links}
		fmt.Println(test_product)
	} //test product
	{
		test_product_type := product_type{super: super{id: -1, name: "molochka", created_at: "22.03.2022", updated_at: "22.03.2022", deleted_at: "22.03.2022"}, image_link: "yandex.ru"}
		fmt.Println(test_product_type)
	} //test product_type
	{
		test_product_type_product := product_type_product{product_type_id: -1, product_id: -1}
		fmt.Println(test_product_type_product)
	} //test product_type_product
	{
		test_store_type_store := store_type_store{store_type_id: -1, store_id: -1}
		fmt.Println(test_store_type_store)
	} //test store_type_store
}

func main() {
	fmt.Println("Hello Golang!")
	test()
}

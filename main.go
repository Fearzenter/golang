package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*func test() {
	{
		testStore := store{super: super{Id: -1, Name: "Vyasoliy_molochnik", CreatedAt: "22.03.2022", UpdatedAt: "22.03.2022", DeletedAt: "22.03.2022"}, ImageLink: "yandex.ru"}
		fmt.Println(testStore)
	} //test store
	{
		testStoreType := storeType{super: super{Id: -1, Name: "productoviy", CreatedAt: "22.03.2022", UpdatedAt: "22.03.2022", DeletedAt: "22.03.2022"}, ImageLink: "yandex.ru"}
		fmt.Println(testStoreType)
	} //test storeType
	{
		links := [...]string{
			"yandex.ru",
			"google.com",
			"vk.com",
			"youtube.com",
		}
		testProduct := product{super: super{Id: -1, Name: "milk", CreatedAt: "22.03.2022", UpdatedAt: "22.03.2022", DeletedAt: "22.03.2022"}, ImageLinks: links}
		fmt.Println(testProduct)
	} //test product
	{
		testProductType := productType{super: super{Id: -1, Name: "molochka", CreatedAt: "22.03.2022", UpdatedAt: "22.03.2022", DeletedAt: "22.03.2022"}, ImageLink: "yandex.ru"}
		fmt.Println(testProductType)
	} //test productType
	{
		testProductTypeProduct := productTypeProduct{ProductTypeId: -1, ProductId: -1}
		fmt.Println(testProductTypeProduct)
	} //test productTypeProduct
	{
		testStoreTypeStore := storeTypeStore{StoreTypeId: -1, StoreId: -1}
		fmt.Println(testStoreTypeStore)
	} //test storeTypeStore
}*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func saveFile(x []byte) {
	err := os.WriteFile("file.json", x, 0644)
	check(err)
}

func loadFile() {
	loaded, err := os.ReadFile("file.json")
	check(err)
	fmt.Println(string(loaded))
}

func main() {
	{
		fmt.Println("Hello!\n")

		menu := "Menu:\n" +
			"1) Save to file\n" +
			"2) Load from file\n" +
			"0) Exit"

	cycle:
		for {
			fmt.Println(menu)
			var answerMenu int
			fmt.Fscan(os.Stdin, &answerMenu)

			switch answerMenu {
			case 1:
				{
					fmt.Println("File saved")
				}
			case 2:
				{
					fmt.Println("File loaded")
				}
			case 0:
				{
					fmt.Println("Goodbye")
					break cycle
				}
			default:
				fmt.Println("Mistake, try again")
			}
		}
		//test()
	} // menu

	storeJsonBefore := &store{super: super{Id: -2, //encoding
		Name:      "testJson",
		CreatedAt: "2022",
		UpdatedAt: "",
		DeletedAt: ""},
		ImageLink: "",
	}
	storeJsonAfter, _ := json.Marshal(storeJsonBefore)
	fmt.Println(storeJsonAfter)

	fmt.Println("----------------------------------------------------------------")

	var dat map[string]interface{} //decoding
	if err := json.Unmarshal(storeJsonAfter, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	saveFile(storeJsonAfter) //test
	loadFile()
}

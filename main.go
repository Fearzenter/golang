package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func saveFile(name string, data []byte) error {
	return os.WriteFile(name, data, 0644)
}

func loadFile(name string) {
	loaded, err := os.ReadFile(name)
	check(err)
	fmt.Println(string(loaded))
}

func main() {
	{
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

		storeJsonBefore := &Store{super: super{Id: -2, //encoding
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

		saveFile("file.json", storeJsonAfter) //test
		loadFile("file.json")
	}

	/*e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))*/
}

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

const url = "http:\\services.explorecalifornia.org/json/tours.php"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n", response)

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Request.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	//fmt.Printf(content) // raw json

	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}

}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	// to check that there are no errors, don't care about return value so put underscore
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() { // while .More returns something from json object
		err := decoder.Decode((&tour)) // pass in memory address of tour object and that will fill in the data
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
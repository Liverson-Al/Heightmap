package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type JsonResponse struct {
	Results []Results `json:"results"`
}
type Results struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Elevation float64 `json:"elevation"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	var x string
	var y string
	fmt.Print("Latitude: ")
	fmt.Fscan(os.Stdin, &x)
	fmt.Print("Longitude: ")
	fmt.Fscan(os.Stdin, &y)

	ans := JsonResponse{}
	url := "https://api.open-elevation.com/api/v1/lookup?locations=" + x + "," + y
	err := getJson(url, &ans)
	if err != nil {
		println(err)
		return
	}
	res := ans.Results[0]
	fmt.Printf("results: %+v\n", res)

}

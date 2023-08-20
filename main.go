package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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
			println(err)
			return
		}
	}(r.Body)

	return json.NewDecoder(r.Body).Decode(target)
}

// получение высоты по x и y
func getElevation(x float64, y float64) (res float64, err error) {
	ans := JsonResponse{}
	url := "https://api.open-elevation.com/api/v1/lookup?locations=" + fmt.Sprintf("%f", x+10.2) + "," + fmt.Sprintf("%f", y+12.5)
	err = getJson(url, &ans) //получаем ответ от сервера
	if err != nil {
		println(err)
		return
	}
	//fmt.Printf("results: %+v\n", &ans)
	res = ans.Results[0].Elevation
	return res, err
}

// str = "10,10|20,20|41.161758,-8.583933"          n = n^2 (dim of arr)
func getElevationArr(str string, n int) (res [][]float64, err error) {
	ans := JsonResponse{}
	url := "https://api.open-elevation.com/api/v1/lookup?locations=" + str
	err = getJson(url, &ans) //получаем ответ от сервера в виде множества Results
	if err != nil {
		println(err)
		return
	}
	element := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//res[ans.Results[element].Latitude][ans.Results[element].Longitude] = ans.Results[element].Elevation  //По-хорошему сделать map с такими данными
			res[i][j] = ans.Results[element].Elevation
			element++
		}
	}
	return res, err
}

// x - Latitude, y - Longitude, stepStr - step, n - arr size
func getJsonStr(x string, y string, stepStr string, n int) string {
	var str string
	step, _ := strconv.ParseFloat(stepStr, 64)
	memoryX, _ := strconv.ParseFloat(x, 64)
	xFloat := memoryX
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			str = str + x + "," + y + "|"
			xFloat += step
			x = fmt.Sprintf("%f", xFloat)
		}
		yFloat, _ := strconv.ParseFloat(y, 64)
		yFloat += step
		y = fmt.Sprintf("%f", yFloat)
		xFloat = memoryX
	}
	str = str[:len(str)-1] //удаляем последний |
	//fmt.Println(str)
	return str
}

func main() {

	//ввод данных через консоль
	var x string
	var y string
	var stepStr string
	var arrSize string
	fmt.Print("Latitude: ")
	fmt.Fscan(os.Stdin, &x)
	fmt.Print("Longitude: ")
	fmt.Fscan(os.Stdin, &y)
	fmt.Print("Step: ")
	fmt.Fscan(os.Stdin, &stepStr)
	fmt.Print("Array size: ")
	fmt.Fscan(os.Stdin, &arrSize)

	n, err := strconv.Atoi(arrSize)
	if err != nil {
		println(err)
		return
	}
	data := make([][]float64, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//От этой штуки отказываюсь, т.к. быстрее получить от сервера большой json, чем посылать тонну запросов
			//res, err := getElevation(float64(i), float64(j)) //получаем значение одной высоты
			str := getJsonStr(x, y, stepStr, n)
			data, err = getElevationArr(str, n)

			if err != nil {
				println(err)
				return
			}
			//data[i] = append(data[i], res)
			fmt.Printf("%+v\t", data[i][j])
		}
		fmt.Printf("\n")
	}

	//fmt.Println("results: %+v\n", data[2][2])
	//fmt.Printf("results: %+v\n", &data)
}

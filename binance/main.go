package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.binance.com/api/v3/avgPrice?symbol=LINKUSDC")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	// jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	// jsonValue, _ := json.Marshal(jsonData)
	// response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	// if err != nil {
	//     fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	//     data, _ := ioutil.ReadAll(response.Body)
	//     fmt.Println(string(data))
	// }

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	//TODO Cambiar la URL para que acceda a una API diferente
	url := "https://bitpay.com/api/rates"
        var rates []map[string]interface{}

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal([]byte(body), &rates)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//printing decoded array values one by one
	for _, rate := range rates {
		 //TODO Corregir los campos del JSON de acuerdo a la API accedida
         fmt.Println("Code:", rate["code"], "Name:", rate["name"], "Rate:", rate["rate"])
	}
}
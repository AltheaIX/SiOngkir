package http

import (
	"SiOngkir/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

func SiCepatOngkir(origin string, destination string, weight string) (models.SiCepatRes, error) {
	data := models.SiCepatRes{}
	client := http.Client{}
	url := "https://content-main-api-production.sicepat.com/public/delivery-fee/fare-non-international"

	payloadMap := map[string]interface{}{
		"origin":      origin,
		"destination": destination,
		"weight":      weight,
		"p":           "0",
		"l":           "0",
		"t":           "0",
	}
	payload, err := json.Marshal(payloadMap)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	res, err := client.Post(url, "application/json", bytes.NewBuffer(payload))
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	data.Weight = weight
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}
	return data, err
}

func AnterAjaOngkir(origin string, destination string) (models.AnterAjaRes, error) {
	data := models.AnterAjaRes{}
	client := http.Client{}
	param := url2.Values{}
	param.Add("origin", origin)
	param.Add("destination", destination)
	url := "https://anteraja.id/api/api/tracking/trackparcel/getRates?" + param.Encode()
	res, err := client.Post(url, "application/json", nil)
	defer res.Body.Close()
	if err != nil {
		return data, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	data.Weight = 1
	if err != nil {
		return data, err
	}
	return data, nil
}

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

func SiCepatOngkir(request *models.RequestSiCepat) (models.SiCepatRes, error) {
	data := models.SiCepatRes{}
	client := http.Client{}
	url := "https://content-main-api-production.sicepat.com/public/delivery-fee/fare-non-international"

	payloadMap := map[string]interface{}{
		"origin":      request.Request.Origin,
		"destination": request.Request.Destination,
		"weight":      request.Weight,
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
	data.Weight = request.Weight
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}
	return data, err
}

func AnterAjaOngkir(request *models.Request) (models.AnterAjaRes, error) {
	data := models.AnterAjaRes{}
	client := http.Client{}
	param := url2.Values{}
	param.Add("origin", request.Origin)
	param.Add("destination", request.Destination)
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

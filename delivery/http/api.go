package http

import (
	"SiOngkir/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AnterAja struct {
	AnterAja models.AnterAjaRes
}

type Sicepat struct {
	Sicepat models.SicepatRes
}

func SiCepatOngkir() (Sicepat, error) {
	data := Sicepat{}
	client := http.Client{}
	url := "https://content-main-api-production.sicepat.com/public/delivery-fee/fare-non-international"
	payload := []byte(`{"origin":"PBL","destination":"PBL10014","weight":"1","p":0,"l":0,"t":0}`)
	res, err := client.Post(url, "application/json", bytes.NewBuffer(payload))
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {

	}
	return data, err
}

func AnterAjaOngkir() (AnterAja, error) {
	data := AnterAja{}
	client := http.Client{}
	url := "https://anteraja.id/api/api/tracking/trackparcel/getRates?origin=35.13.12&destination=32.15.13"
	payload := []byte(`{"client_code":"ACA","origin":"35.13.12","destination":"32.15.13","weight":1}`)
	res, err := client.Post(url, "application/json", bytes.NewBuffer(payload))
	defer res.Body.Close()
	if err != nil {
		return data, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data.AnterAja)
	if err != nil {
		return data, err
	}
	return data, nil
}

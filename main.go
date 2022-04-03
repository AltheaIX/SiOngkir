package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SicepatRes struct {
	Status struct {
		Code        int
		Description string
	}
	Results []struct {
		Service     string
		Description string
		Tariff      int
		MinPrice    int
		UnitPrice   int
		Etd         string
	}
}

type AnterAjaRes struct {
	Status  float32
	Info    string
	Content struct {
		Origin      string
		Destination string
		Services    []struct {
			ProductCode string `json:"product_code"`
			ProductName string `json:"product_name"`
			Etd         string
			Rates       float64
			ImgUrl      string
			Idx         float32
			MsgId       string `json:"msg_id"`
			MsgEn       string `json:"msg_en"`
			InfoId      string `json:"info_id"`
			InfoEn      string `json:"info_en"`
			Enable      bool
		}
	}
}

type AnterAja struct {
	AnterAja AnterAjaRes
}

type Sicepat struct {
	Sicepat SicepatRes
}

func SiCepatOngkir() {
	data := Sicepat{}
	client := http.Client{}
	url := "https://content-main-api-production.sicepat.com/public/delivery-fee/fare-non-international"
	payload := []byte(`{"origin":"PBL","destination":"PBL10014","weight":"1","p":0,"l":0,"t":0}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("==================== SICEPAT ====================")
	for _, each := range data.Sicepat.Results {
		fmt.Printf("Nama Service: %s \n", each.Service)
		fmt.Printf("Estimasi Sampai: %s \n", each.Etd)
		fmt.Printf("Tarif: %d \n", each.Tariff)
		fmt.Println("=================================================")
	}
}

func AnterAjaOngkir() {
	data := AnterAjaRes{}
	client := http.Client{}
	url := "https://anteraja.id/api/api/tracking/trackparcel/getRates?origin=35.13.12&destination=32.15.13"
	payload := []byte(`{"client_code":"ACA","origin":"35.13.12","destination":"32.15.13","weight":1}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)
	fmt.Println("=================== ANTERAJA ====================")
	for _, each := range data.Content.Services {
		tarif := ""
		tarif = fmt.Sprintf("%0.f", each.Rates)
		if !each.Enable {
			tarif = each.MsgId
		}

		fmt.Printf("Nama Service: %s \n", each.ProductName)
		fmt.Printf("Estimasi Sampai: %s \n", each.Etd)
		fmt.Printf("Tarif: %s \n", tarif)
		fmt.Println("=================================================")
	}
}

func main() {
	AnterAjaOngkir()
	SiCepatOngkir()
}

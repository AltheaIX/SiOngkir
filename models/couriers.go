package models

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

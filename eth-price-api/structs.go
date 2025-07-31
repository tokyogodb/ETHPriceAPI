package main

type ETHData struct{
  Crypto      string `json:"symbol"`
  PriceUSDT   string `json:"price"`
}

type Response struct{
	Crypto      string 
    PriceUSDT   float64
	Exchange    string  
	TimeNow     string  
	Volume      float64 
}

type Volume struct{
	Volume string `json:"volume"`
}

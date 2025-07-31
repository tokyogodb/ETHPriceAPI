package main 

import(
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"net/http"
)

func ETHPriceMEXC (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	respAPI, err := http.Get("https://api.mexc.com/api/v3/ticker/price?symbol=ETHUSDT") // Getting data from MEXC API
	if err != nil{
        log.Fatal(err)
	}

	defer respAPI.Body.Close()

	// Decode the received JSON data into the ETHData structure
     var APIData ETHData
	 json.NewDecoder(respAPI.Body).Decode(&APIData) 

	 // Convert Data
	 price, err := strconv.ParseFloat(APIData.PriceUSDT, 64)
	 if err != nil{
		fmt.Println("Conversion Error")
		price = 0
	 }

	 
	volumeData, err := http.Get("https://api.mexc.com/api/v3/ticker/24hr?symbol=ETHUSDT")
	if err != nil{
	   log.Fatal(err)
	}
    
	defer volumeData.Body.Close()

	var APIvolumeData Volume
	json.NewDecoder(volumeData.Body).Decode(&APIvolumeData)

	volume, err := strconv.ParseFloat(APIvolumeData.Volume, 64)
	 if err != nil{
		fmt.Println("Conversion Error")
		volume = 0
	 }

	var resp Response

	resp.Exchange = "MEXC"
	resp.Crypto = APIData.Crypto
	resp.PriceUSDT = price
	resp.Volume = volume
	resp.TimeNow = time.Now().Format(time.RFC1123)
	
    
	json.NewEncoder(w).Encode(resp)
}

func ETHPriceBinance(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  
  respAPI, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT")
  if err != nil{
    log.Fatal(err)
  }
  
  defer respAPI.Body.Close()

  // Decode the received JSON data into the ETHData structure
     var APIData ETHData
   json.NewDecoder(respAPI.Body).Decode(&APIData)

   // Convert Data
   price, err := strconv.ParseFloat(APIData.PriceUSDT, 64)
   if err != nil{
    fmt.Println("Conversion Error")
    price = 0
   }

    
  volumeData, err := http.Get("https://api.binance.com/api/v3/ticker/24hr?symbol=ETHUSDT")
  if err != nil{
     log.Fatal(err)
  }
    
  defer volumeData.Body.Close()

  var APIvolumeData Volume
  json.NewDecoder(volumeData.Body).Decode(&APIvolumeData)

  volume, err := strconv.ParseFloat(APIvolumeData.Volume, 64)
  if err != nil{
    fmt.Println("Conversion Error")
    volume = 0
  }

  var resp Response

  resp.Exchange = "Binance"
  resp.Crypto = APIData.Crypto
  resp.PriceUSDT = price
  resp.Volume = volume
  resp.TimeNow = time.Now().Format(time.RFC1123)
    
  json.NewEncoder(w).Encode(resp)
 }

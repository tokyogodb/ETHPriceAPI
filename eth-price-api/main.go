package main

import (
	"fmt"
	"net/http"
)

 func main(){
    http.HandleFunc("/ETHPriceMEXC", ETHPriceMEXC)
	http.HandleFunc("/ETHPriceBinance", ETHPriceBinance)
	fmt.Println("server is running")
	http.ListenAndServe(":8080", nil)
 }	

 
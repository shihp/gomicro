package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"micro-account-service/model"
)

func main() {

	url := "http://192.168.99.100:6767/accounts/10010"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "8c1f5075-e917-4acc-b8e7-192d7daa4a2f")

	var account model.Account
	for a := 0; a < 1000; a++ {
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		json.Unmarshal([]byte(body), &account)
		fmt.Println(account.ServedBy + " ^ " +  account.Quote.ServedBy)
		//fmt.Println(res)
		//fmt.Println(string(body))
	}
}
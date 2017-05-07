package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"runtime"
)

func main() {

	companies := []string {"googl", "appl", "msft", "bbry", "hpq", "vz", "t", "tmus", "s"}
	start := time.Now()
	fetchedCompanies := 0

	runtime.GOMAXPROCS(4)

	for _, company := range companies {
		go func(company string) {
			var share Share
			resp, err := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote/json?symbol=" + company);

			if err != nil {
				fmt.Println("failed unable to parse");
				return
			}
			err = json.NewDecoder(resp.Body).Decode(&share)
			if err != nil {
				fmt.Println("unable to parse json", err)
			}
			fmt.Println("Symbol", company, "Company", share.Name, ": ", share.LastPrice)
			fetchedCompanies++
		}(company)
	}

	for fetchedCompanies < len(companies) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Println("time taken for above is", elapsed)
}

type Share struct{
	Name string								`json:"Name"`
	LastPrice float32					`json:"LastPrice"`
	Symbol string							`json:"Symbol"`
}

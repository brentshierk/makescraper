package main

import (
	"encoding/json"
	"fmt"
	//"os"
	// "fmt"

	//"io/ioutil"
	"time"

	"github.com/gocolly/colly"
	//"log"
)
type Coin struct{
	Coin_rank  string`json:"Coin_rank"`
	//Coin_image string
	Coin_name string `json:"Coin_name"`
	Coin_price string `json:"Coin_price"`
	Coin_1hr string`json:"Coin_1hr"`
	Coin_24hr string`json:"Coin_24hr"`
	Coin_7d string`json:"Coin_7d"`
}

func main() {
	// fName := "cryptocoinmarketcap.json"
	// file, err := os.Create(fName)
	// if err != nil {
	// 	log.Fatalf("Cannot create file %q: %s\n", fName, err)
	// 	return
	// }
	
	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(e *colly.HTMLElement, ) {
		newCoin := &Coin{
			Coin_rank : e.ChildText(".cmc-table__cell--sort-by__rank"),
			Coin_name : e.ChildText(".cmc-table__column-name"),
			//Coin_image : e.ChildText(".cmc-static-icon"),
			Coin_price : e.ChildText(".cmc-table__cell--sort-by__price"),
			Coin_1hr : e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"),
			Coin_24hr : e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"),
			Coin_7d : e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"),
			
			
		}
		//usign time.sleep to allow for scraper to grab all of the top 20 coins
		time.Sleep(200)
		
		//struct to json 
		
		t,err := json.Marshal(newCoin)
		if err != nil{
			fmt.Printf("error: %s",err)
		}
		fmt.Println(string(t))
		
		
		
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	
}


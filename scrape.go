package main

import (
	"encoding/json"
	"fmt"
	

	"github.com/gocolly/colly"
)
type Coin struct{
	Coin_rank  string
	Coin_image string
	Coin_name string 
	Coin_price string 
	Coin_1hr string
	Coin_24hr string
	Coin_7d string
}
type Counter struct{
	counter int
}
func (self Counter) currentValue() int {
	return self.counter
}
func (self *Counter) increment() {
	self.counter++
	
}
// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	
pi:= Counter{0}
pi.increment()
	
	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(e *colly.HTMLElement, ) {
		newCoin := &Coin{
			Coin_rank : e.ChildText(".cmc-table__cell--sort-by__rank"),
			Coin_name : e.ChildText(".cmc-table__column-name"),
			Coin_image : e.ChildAttr(".cmc-static-icon-1","img"),
			Coin_price : e.ChildText(".cmc-table__cell--sort-by__price"),
			Coin_1hr : e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"),
			Coin_24hr : e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"),
			Coin_7d : e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"),
			
			
		}
		// newCoin.increment()
		t,err := json.Marshal(newCoin)
		if err != nil{
			fmt.Printf("error: %s",err)
		}
		fmt.Println(string(t))
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	
}

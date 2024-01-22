package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type PokemonProduct struct { 
	url,name string 
} 
 
func main() { 
	// initializing the slice of structs to store the data to scrape 
	// Slice of struct is like array of json in javascript [{},{},{},...,{}]
	// So vaariable pokemonProducts is something like [{url:xxx,name:333},{},...{}]
	var pokemonProducts []PokemonProduct 
 
	// creating a new Colly instance 
	c := colly.NewCollector() 
 
	// scraping logic 
	c.OnHTML("h3.wb-break-all", func(e *colly.HTMLElement) { 
		pokemonProduct := PokemonProduct{} 
		pokemonProduct.url = e.ChildAttr("a", "href") 
		pokemonProduct.name = e.ChildText("a") 
		pokemonProducts = append(pokemonProducts, pokemonProduct) 
	}) 
	// visiting the target page 
	c.Visit("https://github.com/fojanb?tab=repositories") 
	// opening the CSV file 
	file, err := os.Create("products.csv") 
	if err != nil { 
		fmt.Print("Failed to create output CSV file", err)
		return
	} 
	defer file.Close() 
 
	// initializing a file writer 
	writer := csv.NewWriter(file) 
 
	// writing the CSV headers 
	headers := []string{ 
		"url", 
		"name", 
	} 
	writer.Write(headers) 
 
	// writing each Pokemon product as a CSV row 
	for _, pokemonProduct := range pokemonProducts { 
		// converting a PokemonProduct to an array of strings 
		record := []string{ 
			pokemonProduct.url, 
			pokemonProduct.name, 
		} 
 
		// adding a CSV record to the output file 
		writer.Write(record) 
	} 
	defer writer.Flush() 
}

	
/*These functions are executed in the following order:

OnRequest(): Called before performing an HTTP request with Visit().
OnError(): Called if an error occurred during the HTTP request.
OnResponse(): Called after receiving a response from the server.
OnHTML(): Called right after OnResponse() if the received content is HTML.
OnScraped(): Called after all OnHTML() callback executions.
---------------------------------------------------------------------------

Here, note that a target li HTML element has the .product class and stores:

An a element with the product URL.
An img element with the product image.
An h2 element with the product name.
A .price element with product pric
*/
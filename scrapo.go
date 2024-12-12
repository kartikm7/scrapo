package main

import (
	"fmt"

	"github.com/kartikm7/scrapo/parser"
	"github.com/kartikm7/scrapo/scraper"
)

func main() {
	if html, err := scraper.Scraper("https://www.llocal.in"); err == nil {
		if cleaned, err := parser.Parser(html); err == nil {
			fmt.Println(cleaned)
		}
	} else {
		fmt.Println("Dude some error occured")
		fmt.Println(err)
	}
}

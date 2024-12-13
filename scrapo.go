package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kartikm7/scrapo/parser"
	"github.com/kartikm7/scrapo/scraper"
)

// type for the website array/slice idk the difference yet
type website struct {
	name string
	url  string
}

// wait group which acts basically like a counter
var wg = sync.WaitGroup{}

// mutual exclusion
var mutex = sync.Mutex{}

// this is mock data to test go routines
var websites []website = []website{
	{"LLocal", "https://www.llocal.in"},
	{"My GitHub", "https://www.github.com/kartikm7"},
	{"DocNext", "https://docnext.llocal.in"},
	{"Ono", "https://ono.llocal.in"},
}

var result = []string{}

func main() {
	startingTime := time.Now()
	for _, val := range websites {

		// increment the counter
		wg.Add(1)
		go func() {
			if data, err := scraper.Scraper(val.url); err == nil {
				// I don't quite care about the scraped data as much as I do with the order of processes
				if data, err := parser.Parser(data); err == nil {
					mutex.Lock()
					fmt.Println(data)
					result = append(result, val.name)
					mutex.Unlock()
				}
			}
			// decrements the counter by 1
			wg.Done()
		}()
	}
	// it waits here till the counter hits 0!
	wg.Wait()
	fmt.Printf("The results: %v\n", result)
	fmt.Println("Time since execution: ", time.Since(startingTime))
}

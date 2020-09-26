package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func req(symbol string) {
	res, err := http.Get("http://localhost:3000/stocks?symbol=" + symbol)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", data)
}

func main() {
	symbols := []string{"PEAR", "SQRL", "NULL", "ZVZZT"}
	for i := 0; i < 1000; i++ {
		symbol := symbols[rand.Intn(len(symbols))]
		req(symbol)
	}
}

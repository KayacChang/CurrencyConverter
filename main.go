package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	base, target := strings.ToUpper(args[0]), strings.ToUpper(args[1])

	money, err := strconv.ParseFloat(args[2], 64)

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Get("https://api.exchangerate-api.com/v4/latest/" + base)

	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		log.Fatal(err)
	}

	_rates, ok := data["rates"].(map[string]interface{})

	if !ok {
		log.Fatal(ok)
	}

	rate, ok := _rates[target].(float64)

	if !ok {
		log.Fatal(ok)
	}

	fmt.Println(rate * money)
}

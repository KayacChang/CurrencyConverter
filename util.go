package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func print(obj interface{}) {
	text, err := json.MarshalIndent(obj, "", "	")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(text))
}

package main

import (
	"encoding/json"
	"fmt"

	"github.com/barokurniawan/goktparser/src"
)

func main() {
	Ktparser := src.NewKtparser()
	output := Ktparser.ParseNIK("3204110609970001")

	jsonOutput, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonOutput))
}

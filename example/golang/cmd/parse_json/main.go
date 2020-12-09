package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"src/model"
)

func main() {
	bytes, err := ioutil.ReadFile("./data/2e22de6b770e4166be301490f6ffd420.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	v := new(model.BlockMap)
	err = json.Unmarshal(bytes, v)
	if err != nil {
		log.Fatal(err)
		return
	}
	blockType := (*v)["77135bf3-fa11-4c19-8397-69ee133ca1ce"]
	data, err := json.MarshalIndent(blockType, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(data))
}

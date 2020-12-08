package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"src/scheme"
)

func main() {
	bytes, err := ioutil.ReadFile("./data/2e22de6b770e4166be301490f6ffd420.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	v := new(scheme.BlockMapType)
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

	like := Like{
		Name: "张三",
		Age:  22,
	}

	object := TestObject{
		Like:        like,
		ID:          12,
		CreatedTime: 1996,
	}
	ov, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(ov))
	oc := new(TestObject)
	err = json.Unmarshal(ov, oc)
	if err != nil {
		log.Fatal(err)
		return
	}
}

type Like struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type TestObject struct {
	Like        `json:",inline"`
	ID          int64
	CreatedTime int64
}

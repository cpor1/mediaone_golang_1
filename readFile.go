package main

import (
	"encoding/json"
	"io/ioutil"

	"gitlab.com/mediaone/test/common"
)

func readFile() []Serve {
	file, err := ioutil.ReadFile("serve.json")
	common.Check(err)
	serves := make([]Serve, 2)
	json.Unmarshal([]byte(file), &serves)
	return serves
}

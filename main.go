package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	jsonPath, templatePath := getFlags()
	outputFlags(jsonPath, templatePath)
	json, err := readJson(jsonPath)

	if err != nil {
		fmt.Printf("read json error: %s", err)
		return
	}

	fmt.Println(json)
}

func outputFlags(jsonPath string, templatePath string) {
	fmt.Printf("json path: %s\n", jsonPath)
	fmt.Printf("template path: %s\n", templatePath)
}

func getFlags() (jsonPath, templatePath string) {
	flag.StringVar(&jsonPath, "json-path", "./json.json", "json file path")
	flag.StringVar(&templatePath, "template-path", "./template.template", "template file path")
	flag.Parse()

	return
}

func readJson(jsonPath string) (json string, err error) {
	var data []byte
	data, err = ioutil.ReadFile(jsonPath)

	if err != nil {
		return
	}

	json = string(data)

	foo := parseJson(data)
	fmt.Printf("%v", foo)

	return
}

func parseJson(data []byte) (result map[string]interface{}) {
	var f interface{}
	err := json.Unmarshal(data, &f)

	if err == nil {
		result = f.(map[string]interface{})
	}

	return result
}

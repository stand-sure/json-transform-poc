package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

// entry point
func main() {
	jsonPath, templatePath := getFlags()
	data, err := readJson(jsonPath)

	if err != nil {
		fmt.Printf("read json error: %s", err)
		return
	}

	t := readTemplate(templatePath)
	t.Execute(os.Stdout, data)
}

// getFlags parses the command-line argumentss
func getFlags() (jsonPath, templatePath string) {
	flag.StringVar(&jsonPath, "json-path", "./sample.json", "json file path")
	flag.StringVar(&templatePath, "template-path", "./default.template", "template file path")
	flag.Parse()

	return
}

// readJson reads the file at jsonPath
// and returns a map[string]interface and an error
func readJson(jsonPath string) (raw map[string]interface{}, err error) {
	var data []byte
	data, err = ioutil.ReadFile(jsonPath)

	if err != nil {
		return
	}

	raw = parseJson(data)

	return
}

// parseJson parses a byte slice to a map with string keys
func parseJson(data []byte) (result map[string]interface{}) {
	var f interface{}
	err := json.Unmarshal(data, &f)

	if err == nil {
		result = f.(map[string]interface{})
	}

	return result
}

// readTemplate reads the template found at the path in templatePath
// and returns the *template.Template
// IMPORTANT: if the template is invalid, the application will PANIC
func readTemplate(templatePath string) (t *template.Template) {
	t = template.Must(template.ParseFiles(templatePath))
	return
}

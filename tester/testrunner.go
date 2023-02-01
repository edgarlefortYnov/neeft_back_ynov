package main

import (
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"io"
	"io/ioutil"
	"net/http"
	"tester/models"
)

func RunByName(name string) bool {
	return Run(fmt.Sprintf("../tests/definitions/%s.json", name))
}

func Run(path string) bool {
	var def models.TestDefinition
	var client http.Client

	// TODO: Replace ioutil.ReadFile()
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}

	err = json.Unmarshal(data, &def)
	if err != nil {
		return false
	}

	fmt.Printf("[TEST] %s: ", color.Ize(color.Yellow, def.Name))

	req, err := http.NewRequest(def.Method, def.Url, nil)
	response, err := client.Do(req)
	if err != nil {
		return false
	}

	for _, header := range def.Headers {
		req.Header.Add(header.Name, header.Value)
	}

	body, err := io.ReadAll(response.Body)

	return true
}

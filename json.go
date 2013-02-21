package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func jsonFromFile(file string) (jsonData map[string]interface{}, err error) {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		return make(map[string]interface{}), err
	}

	var data interface{}
	err = json.Unmarshal(text, &data)
	if err != nil {
		return make(map[string]interface{}), err
	}

	return data.(map[string]interface{}), nil
}

func valueToString(value interface{}) (text string, err error) {
	if value == nil && *printNulls == false {
		return "", nil
	}

	textBytes, err := json.Marshal(value)
	return string(textBytes), err
}

func get(data *map[string]interface{}, attributeChain []string) (value string, err error) {
	attribute := attributeChain[0]

	if len(attributeChain) == 1 {
		return valueToString((*data)[attribute])
	}

	subdata := (*data)[attribute].(map[string]interface{})
	return get(&subdata, attributeChain[1:])
}

func getValues(data *map[string]interface{}, attributeChains []string) (values []string, err error) {
	values = make([]string, len(attributeChains))
	for i, attributeChain := range attributeChains {
		value, err := get(data, strings.Split(attributeChain, "."))
		if err != nil {
			return values, nil
		}
		values[i] = value
	}
	return values, nil
}

package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadLimitedJsonFile2Words(jsonFile *os.File) ([]Word, error) {

	outputWords := make([]Word, 0)
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(byteValue), &outputWords)
	return outputWords, nil
}

func Decode2WordsInLimit(r io.Reader) ([]Word, error) {
	dec := json.NewDecoder(r)

	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	if t != json.Delim('{') {
		return nil, fmt.Errorf("expected {, got %v", t)
	}

	outputWords := make([]Word, 0)
	index := 1
	// While there are more tokens in the JSON stream...
	for dec.More() {

		// Read the key.
		t, err := dec.Token()
		if err != nil {
			return nil, err
		}
		key := t.(string) // type assert token to string.

		// Decode the value.
		var value string
		if err := dec.Decode(&value); err != nil {
			return nil, err
		}

		var word = Word{
			Index:      index,
			Word:       key,
			Definition: value,
		}
		outputWords = append(outputWords, word)
		index++
	}
	return outputWords, nil
}

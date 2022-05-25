package internal

import (
	"encoding/json"
	"fmt"
	"io"
)

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
			Word:       key,
			Definition: value,
		}
		outputWords = append(outputWords, word)
	}
	return outputWords, nil
}

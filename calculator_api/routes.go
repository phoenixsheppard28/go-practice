package main

import (
	"encoding/json"
	"net/http"
)

func Route1(w http.ResponseWriter, _ *http.Request) {
	encoder := json.NewEncoder(w)

	data := map[string]interface{}{
		"Hello": 2,
		"World": 5,
	}
	encoder.Encode(data)

}

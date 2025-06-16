package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Input struct {
	Num1     *int   `json:"num1" validate:"required"`
	Num2     *int   `json:"num2" validate:"required"`
	Operator string `json:"operator" validate:"required"`
}

func handle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Error: Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Error: Invalid json payload", 400)
		return
	}
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		http.Error(w, "Error invalid input", 400)
		return
	}
	num1 := *input.Num1
	num2 := *input.Num2
	oper := input.Operator
	encoder := json.NewEncoder(w)

	var result int
	switch oper {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		if num2 == 0 {
			http.Error(w, "Error: Divide by zero", 400)
			return
		}
		result = num1 / num2
	case "*":
		result = num1 * num2
	case "%":
		if num2 == 0 {
			http.Error(w, "Error: Mod by zero", 400)
			return
		}
		result = num1 % num2
	default:
		http.Error(w, "Error: Invalid operation, only allow + - * / %", 400)
		return
	}

	err = encoder.Encode(result)
	if err != nil {
		http.Error(w, "Error: failed to encode response", 500)
	}
}

func main() {
	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

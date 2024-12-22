package application

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"Calculation-Service-Yandex/pkg/calc"
)

type Request struct {
	Expression string `json:"expression" `
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func sendResponse(w http.ResponseWriter, response Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func CalculationHandler(w http.ResponseWriter, r *http.Request, calcFunc func(string) (float64, error)) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %v\n", err)
		response := Response{Error: "Invalid JSON format"}
		sendResponse(w, response, http.StatusBadRequest)
		return
	}

	result, err := calcFunc(request.Expression) // Теперь используем calcFunc
	if err != nil {
		log.Printf("Error calculating expression: %v\n", err)
		var response Response
		status := http.StatusInternalServerError

		switch {
		case errors.Is(err, calc.ErrInvalidExpression):
			response.Error = "Expression is not valid"
			status = http.StatusUnprocessableEntity
		case errors.Is(err, calc.ErrDivisionByZero):
			response.Error = "Division by zero"
			status = http.StatusUnprocessableEntity
		default:
			response.Error = "Internal server error"
		}

		sendResponse(w, response, status)
		return
	}
	response := Response{Result: result}
	sendResponse(w, response, http.StatusOK)
}

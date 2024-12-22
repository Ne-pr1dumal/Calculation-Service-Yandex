package main

import (
	"fmt"
	"log"
	"net/http"

	"Calculation-Service-Yandex/internal/application"
	"Calculation-Service-Yandex/pkg/calc"
)

func main() {
	// Задаем порт по умолчанию
	port := 8080
	// Создаем новый хендлер
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		application.CalculationHandler(w, r, calc.Calc)
	})
	// Запускаем сервер
	fmt.Printf("server listen at port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

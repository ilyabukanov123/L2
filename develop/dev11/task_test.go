package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestEventsForDay(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/events_for_day")
	if err != nil {
		log.Fatalf("Произошла ошибка при отправке GET запроса на http:/localhost:8080/events_for_month: %s\n", err)
	}
	fmt.Println(resp)
}

func TestCreateEvent(t *testing.T) {
	message := map[string]interface{}{
		"event_id":    1,
		"user_id":     1,
		"name":        "Событие №1",
		"description": "Очень важное событие 1",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Произошла ошибка при конвертации map в JSON %s\n", err)
	}

	resp, err := http.Post("http://localhost:8080/create_event", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalf("Произошла ошибка при отправке GET запроса на http:/localhost:8080/create_event: %s\n", err)
	}
	fmt.Println(resp)
}

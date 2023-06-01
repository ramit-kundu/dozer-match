package catscraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GptParser(html string) {
	url := "https://api.openai.com/v1/chat/completions"

	// Define the request payload
	payload := struct {
		Model    string `json:"model"`
		Messages []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
	}{
		Model: "gpt-3.5-turbo",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: "Hello!",
			},
		},
	}

	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")

	// Send the request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Read the response body
	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	// Print the response
	fmt.Println(result)
}

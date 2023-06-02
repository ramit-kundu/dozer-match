package catscraper

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/kundu-ramit/dozer_match/domain/entity"
)

func GptParser(ctx context.Context, html string) (*entity.BullDozer, error) {
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
				Role: "user",
				Content: `From the below html snippet find out 
				○ Make ( Brand eg Caterpillar)
				○ Model ( Model No )
				○ Picture (check for link)
				○ Category (i.e., Small Dozer, Medium Dozer, Large Dozer, Wheel Dozer)
				○ Engine HP 
				○ Operating Weight (in lbs or kgs)
				and show it in json format. Only respond with the json.
				Ex : {
					"Make": "Caterpillar",
					"Model": "D1",
					"Picture": "https://s7d2.scene7.com/is/image/Caterpillar/CM20200423-2bfc1-01ee1?$cc-s$",
					"Category": "Small Dozers",
					"Engine HP": "80 HP",
					"Operating Weight": "17855 lb"
				  }
				  
				  html is : ` + html,
			},
		},
	}

	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")

	// Send the request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read the response body
	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	apiRes := result["choices"].([]map[string]interface{})[0]["message"].(map[string]map[string]interface{})["content"]

	return &entity.BullDozer{
		Make:            apiRes["Make"].(string),
		Model:           apiRes["Model"].(string),
		Picture:         apiRes["Picture"].(string),
		Category:        apiRes["Category"].(string),
		EngineHP:        apiRes["EngineHP"].(string),
		OperatingWeight: apiRes["OperatingWeight"].(int64),
	}, nil

}

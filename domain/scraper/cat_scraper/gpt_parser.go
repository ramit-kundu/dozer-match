package catscraper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
				Role: "system",
				Content: `You are a helpful assistant that inputs a html snippet 
				which is scraped from a website of bulldozers and from that html snippet 
				finds out 
				○ Make ( Brand eg Caterpillar)
				○ Model ( Model No )
				○ Picture (check for link)
				○ Category (i.e., Small Dozer, Medium Dozer, Large Dozer, Wheel Dozer)
				○ Engine HP 
				○ Operating Weight (in lbs or kgs)
				of the bulldozer. You will only respond with a json in this format
				{
					"Make": "Caterpillar",
					"Model": "D1",
					"Picture": "https://s7d2.scene7.com/is/image/Caterpillar/CM20200423-2bfc1-01ee1?$cc-s$",
					"Category": "Small Dozers",
					"EngineHP": "80 HP",
					"OperatingWeight": "17855 lb"
				  } If youre unable to find any component in json plz put empty string and not null
			 `,
			},
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
					"EngineHP": "80 HP",
					"OperatingWeight": "17855 lb"
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
	choices := result["choices"].([]interface{})
	fmt.Println(choices...)
	message := choices[0].(map[string]interface{})["message"]
	fmt.Println(message)
	content := message.(map[string]interface{})["content"].(string)
	fmt.Println(content)
	apiRes := make(map[string]interface{})
	json.Unmarshal([]byte(content), &apiRes)

	opWt, _ := strconv.ParseInt(apiRes["OperatingWeight"].(string), 10, 64)

	return &entity.BullDozer{
		Make:            apiRes["Make"].(string),
		Model:           apiRes["Model"].(string),
		Picture:         apiRes["Picture"].(string),
		Category:        apiRes["Category"].(string),
		EngineHP:        apiRes["EngineHP"].(string),
		OperatingWeight: opWt,
	}, nil

}

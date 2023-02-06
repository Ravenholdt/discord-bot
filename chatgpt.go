package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
result, err := AskChatGPT("What is the meaning of life?")
if err != nil {
	println("Error:", err)
	return
}

println("Result:", result)
*/

const openaiURL = "https://api.openai.com/v1/engines/text-davinci/jobs"

// QuestionResponse is the struct used to parse the response from the OpenAI API
type QuestionResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// AskChatGPT sends a question to the OpenAI API and returns the response as a string
func AskChatGPT(question string) (string, error) {
	// Create a new HTTP request
	requestBody := map[string]interface{}{
		"prompt":      question,
		"max_tokens":  100,
		"temperature": 0.5,
	}

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestBytes))
	if err != nil {
		return "", err
	}

	// Add the API key to the request headers
	req.Header.Add("Authorization", "Bearer <YOUR_OPENAI_API_KEY>")
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Read the response
	defer resp.Body.Close()
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response
	var response QuestionResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return "", err
	}

	return response.Message, nil
}

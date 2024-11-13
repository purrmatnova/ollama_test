// Example how to use your own model via ollama
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Role    string   `json:"role"`
	Content string   `json:"content"`
	Images  []string `json:"images"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func main() {

	imagePath := "cat.jpg"

	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Error reading image:", err)
		return
	}

	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	requestBody := Request{
		Model: "purrmatnova/cats-vs-dogs",
		Messages: []Message{
			{
				Images: []string{encodedImage},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(responseBody))
}

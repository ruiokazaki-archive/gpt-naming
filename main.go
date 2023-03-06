package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	apiKey := getApiKey()
	answers := askRequest()
	result := sendRequest(apiKey, answers)

	fmt.Println(result.Choices[0].Text)
}

func createFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v", err)
		os.Exit(1)
	}

	dirPath := filepath.Join(homeDir, ".config", "gpt-naming")
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Printf("Error creating directory: %v", err)
		os.Exit(1)
	}

	filePath := filepath.Join(dirPath, "api_key")

	return filePath
}

func getApiKey() string {
	filePath := createFilePath()
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		tokenQuestion := []*survey.Question{
			{
				Name: "token",
				Prompt: &survey.Password{
					Message: "Enter an OpenAi-Token:",
				},
				Validate: survey.Required,
			},
		}

		answer := TokenAnswer{}

		err := survey.Ask(tokenQuestion, &answer)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file: %v", err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.WriteString(answer.Token)
		if err != nil {
			fmt.Printf("Error writing to file: %v", err)
			os.Exit(1)
		}

		fmt.Println("api_key file created at", filePath)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	lines := strings.Split(string(content), "\n")

	OPENAI_API_KEY := lines[0]

	return OPENAI_API_KEY
}

func askRequest() MainAnswers {
	questions := []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Choose a type:",
				Options: []string{"argument", "class", "constant", "delegate", "enum", "event", "exception", "function", "interface", "method", "namespace", "property", "struct", "type", "variables"},
				Default: "function",
			},
			Validate: survey.Required,
		},
		{
			Name: "overview",
			Prompt: &survey.Input{
				Message: "Enter an overview:",
			},
			Validate: survey.Required,
		},
	}

	answers := MainAnswers{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return answers
}

func sendRequest(apiKey string, answers MainAnswers) ApiResponse {
	prompt := fmt.Sprintf(
		`# Request\nThink of a %s name to use in programming.\n# Function Summary\n%s\n# Condition\n- Output in lower camel case\n- Briefly explain the reason for naming\n- Output the reason for naming in the language you are outlining\n- Output up to 5 outputs according to the format\n# Output format\nindex. [naming]: reason`, answers.Type, answers.Overview)

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/completions",
		bytes.NewBuffer(
			[]byte(`{
				"model": "text-davinci-003",
				"prompt": "`+prompt+`",
				"temperature": 0,
				"max_tokens": 300
			}`)),
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonResult := ApiResponse{}

	if err := json.Unmarshal(result, &jsonResult); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if jsonResult.Error.Message != "" {
		fmt.Println(jsonResult.Error.Message)
		os.Exit(1)
	}

	return jsonResult
}

type MainAnswers struct {
	Type     string
	Overview string
}
type TokenAnswer struct {
	Token string
}
type ApiResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     string `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param"`
		Code    string `json:"code"`
	} `json:"error"`
}

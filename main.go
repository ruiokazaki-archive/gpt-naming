package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

var questions = []*survey.Question{
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "Choose a type:",
			Options: []string{"関数", "変数"},
			Default: "関数",
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

func main() {
	answers := struct {
		Type     string
		Overview string
	}{}

	err := survey.Ask(questions, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prompt := fmt.Sprintf(
		`プログラミングで使用する%s名を考えてください。\n概要\n%s\n条件\n・lower camel caseで出力する\n・フォーマットに従い最大5つ出力する\nフォーマット\n[名前]: 特徴`, answers.Type, answers.Overview)

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/completions",
		bytes.NewBuffer(
			[]byte(`{
				"model": "text-davinci-003",
				"prompt": "`+prompt+`",
				"temperature": 0,
				"max_tokens": 200
			}`)),
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	resultStruct := struct {
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
	}{}

	if err := json.Unmarshal(result, &resultStruct); err != nil {
		fmt.Println(err)
		return
	} else if resultStruct.Error.Message != "" {
		fmt.Println(resultStruct.Error.Message)
		return
	}

	fmt.Println(resultStruct.Choices[0].Text)

}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AlecAivazis/survey/v2"
)

var questions = []*survey.Question{
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "Choose a type:",
			Options: []string{"function", "variable"},
			Default: "function",
		},
		Validate: survey.Required,
	},
	{
		Name: "overview",
		Prompt: &survey.Multiline{
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

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/completions",
		bytes.NewBuffer(
			[]byte(`{
				"model": "text-davinci-003",
				"prompt": "Say this is a test",
				"temperature": 0,
				"max_tokens": 70
			}`)),
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Add("Authorization", "Bearer ")
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

	fmt.Println(string(result))

}

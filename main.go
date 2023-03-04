package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var questions = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "What is your name?",
		},
		Validate: survey.Required,
	},
}

func main() {
	answers := struct {
		Name string
	}{}

	println("Hello, World!")

	err := survey.Ask(questions, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Hello, %s!", answers.Name)
}
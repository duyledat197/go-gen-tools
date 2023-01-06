package models

import "github.com/manifoldco/promptui"

type CliStepType string

var (
	PROMPT CliStepType = "PROMPT"
	SELECT CliStepType = "SELECT"
)

type CliStep struct {
	Name   string
	Type   CliStepType
	Prompt promptui.Prompt
	Select promptui.Select
	Val    string
}

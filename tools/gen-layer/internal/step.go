package internal

import (
	"fmt"

	"github.com/duyledat197/go-gen-tools/tools/gen-layer/models"
	"github.com/manifoldco/promptui"
)

var (
	Layers    = []string{"all", "delivery", "service", "repository", "cucumber"}
	Methods   = []string{"all", "create", "update", "delete", "list", "retrieve"}
	Databases = []string{"mongo", "postgres", "inmem"}

	LayerMap = map[string]string{
		Layers[1]: "deliveries/grpc",
		Layers[2]: "services",
		Layers[3]: "repositories",
	}
)

var Steps = []*models.CliStep{
	{
		Name: "Choose Name want to generate for CRUD",
		Type: models.PROMPT,
		Prompt: promptui.Prompt{
			Label: "Name",
			Templates: &promptui.PromptTemplates{
				Success: fmt.Sprintf("%s {{ . | green }}%s ", promptui.IconGood, promptui.Styler(promptui.FGGreen)(":")),
				Valid:   fmt.Sprintf("{{ . | blue }}%s ", promptui.Styler(promptui.FGBlue)(":")),
				Invalid: fmt.Sprintf("{{ . | blue }}%s ", promptui.Styler(promptui.FGBlue)(":")),
			},
			Validate: func(str string) error {
				return nil
			},
		},
	},
	{
		Name: "step 2",
		Type: models.SELECT,
		Select: promptui.Select{
			Label: "Choose layer want to generate for CRUD",
			Items: Layers,
			Templates: &promptui.SelectTemplates{
				Active:   fmt.Sprintf("%s {{ . | underline | green }}", promptui.IconSelect),
				Label:    fmt.Sprintf("%s {{ . | blue }}: ", promptui.IconInitial),
				Selected: fmt.Sprintf("%s {{ . | white }}", promptui.IconGood+promptui.Styler(promptui.FGGreen)(" Layer name: ")),
			},
		},
	},

	{
		Name: "step 3",
		Type: models.SELECT,
		Select: promptui.Select{
			Label: "Choose method want to generate",
			Items: Methods,
			Templates: &promptui.SelectTemplates{
				Active:   fmt.Sprintf("%s {{ . | underline | green }}", promptui.IconSelect),
				Label:    fmt.Sprintf("%s {{ . | blue }}: ", promptui.IconInitial),
				Selected: fmt.Sprintf("%s {{ . | white }}", promptui.IconGood+promptui.Styler(promptui.FGGreen)(" Method name: ")),
			},
		},
	},

	{
		Name: "step 4",
		Type: models.SELECT,
		Select: promptui.Select{
			Label: "Choose database using",
			Items: Databases,
			Templates: &promptui.SelectTemplates{
				Active:   fmt.Sprintf("%s {{ . | underline | green }}", promptui.IconSelect),
				Label:    fmt.Sprintf("%s {{ . | blue }}: ", promptui.IconInitial),
				Selected: fmt.Sprintf("%s {{ . | white }}", promptui.IconGood+promptui.Styler(promptui.FGGreen)(" Method name: ")),
			},
		},
	},
}

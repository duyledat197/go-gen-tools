package internal

import (
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/duyledat197/interview-hao/cmd/gen-layer/models"
	"github.com/iancoleman/strcase"
)

func Run() {
	var err error
	for _, step := range Steps {
		switch step.Type {
		case models.PROMPT:
			step.Val, err = step.Prompt.Run()
			if err != nil {
				panic(err)
			}
		case models.SELECT:
			_, step.Val, err = step.Select.Run()
			if err != nil {
				panic(err)
			}
		}
	}
	name := Steps[0].Val
	layer := Steps[1].Val
	// method := Steps[2].Val

	templateModel := &models.Template{
		CamelCase:  strcase.ToLowerCamel(name),
		PascalCase: strcase.ToCamel(name),
	}

	baseDir, _ := os.Getwd()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	pkgDir := path.Dir(filename)

	var layers []string
	if layer == Layers[0] {
		layers = Layers
	} else {
		layers = []string{layer}
	}
	for _, l := range layers {
		if l == Layers[0] {
			continue
		}
		file, err := os.Create(path.Join(baseDir, "internal", LayerMap[l], strcase.ToKebab(name)+".go"))
		if err != nil {
			panic(err)
		}
		templatePath := path.Join(pkgDir, "..", "templates", l+".tpl")
		type Name struct {
			Name string
		}
		tmpl := template.
			Must(template.
				ParseFiles(templatePath))
		if err := tmpl.Execute(file, templateModel); err != nil {
			panic(err)
		}
	}
}

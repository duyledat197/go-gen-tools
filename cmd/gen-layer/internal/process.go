package internal

import (
	"log"
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
	file, err := os.Create(path.Join(baseDir, "internal", LayerMap[layer], "grpc", strcase.ToKebab(name)+".go"))
	if err != nil {
		panic(err)
	}
	switch layer {
	case Layers[1]:
		// templatePath := path.Join(pkgDir, "..", "templates", layer+".tpl")
		templatePath := path.Join(pkgDir, "..", "templates", layer+".tpl")

		log.Println("templatePath", templatePath)
		b, err := os.ReadFile(templatePath)
		if err != nil {
			panic(err)
		}
		tmpl := template.
			Must(template.New(layer).
				Parse(string(b)))
			// ParseFiles(templatePath))
		log.Println(tmpl.DefinedTemplates())

		// if err := tmpl.Execute(os.Stdout, templateModel); err != nil {
		// 	log.Fatalf("execution failed: %s", err)
		// }
		// return
		if err := tmpl.Execute(file, templateModel); err != nil {
			panic(err)
		}
	default:
	}
}

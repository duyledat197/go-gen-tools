package internal

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/duyledat197/interview-hao/cmd/gen-layer/models"
	"github.com/duyledat197/interview-hao/utils/moduleutils"
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
	method := Steps[2].Val

	templateModel := &models.Template{
		CamelCase:  strcase.ToLowerCamel(name),
		PascalCase: strcase.ToCamel(name),
		Module:     moduleutils.GetModuleName(),
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

	var methods []string
	if method == Methods[0] {
		methods = Methods
	} else {
		methods = []string{method}
	}
	for _, l := range layers {
		if l == Layers[0] {
			continue
		}
		file, err := os.Create(path.Join(baseDir, "internal", LayerMap[l], strcase.ToKebab(name)+".go"))
		if err != nil {
			panic(err)
		}
		paths := []string{
			path.Join(pkgDir, "..", "templates", l, "default.tpl"), // root path
		}
		p := path.Join(pkgDir, "..", "templates", l, "default.tpl")
		paths = append(paths, p)
		// for _, m := range methods {
		// 	if m == Methods[0] {
		// 		continue
		// 	}
		// 	p := path.Join(pkgDir, "..", "templates", l, m+".tpl")
		// 	paths = append(paths, p)
		// }
		fmt.Println(methods)

		tmpl := template.
			Must(template.
				ParseFiles(paths...))

		if err := tmpl.ExecuteTemplate(file, "default", templateModel); err != nil {
			panic(err)
		}
		// for _, m := range methods {
		// 	if m == Methods[0] {
		// 		continue
		// 	}
		// 	if err := tmpl.ExecuteTemplate(file, m, templateModel); err != nil {
		// 		panic(err)
		// 	}
		// }
	}
}

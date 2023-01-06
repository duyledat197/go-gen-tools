package internal

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/duyledat197/go-gen-tools/cmd/gen-layer/models"
	"github.com/duyledat197/go-gen-tools/cmd/gen-layer/utils"
	"github.com/duyledat197/go-gen-tools/features"
	"github.com/duyledat197/go-gen-tools/utils/pathutils"

	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

func Run() {
	var err error
	for i, step := range Steps {
		layer := Steps[1].Val
		if i == 3 && !(layer == Layers[0] || layer == Layers[3]) {
			continue
		}
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
	database := Steps[3].Val

	baseDir, _ := os.Getwd()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	pkgDir := path.Dir(filename)

	// get layers
	var layers []string
	if layer == Layers[0] {
		layers = Layers
	} else {
		layers = []string{layer}
	}

	if database != "" {
		layers = append(layers, database)
	}

	// get methods
	var methods []string
	if method == Methods[0] {
		methods = Methods
	} else {
		methods = []string{method}
	}

	templateModel := &models.Template{
		CamelCase:  strcase.ToLowerCamel(name),
		PascalCase: strcase.ToCamel(name),
		Module:     pathutils.GetModuleName(),
		IsCreate:   slices.Contains(methods, Methods[1]),
		IsUpdate:   slices.Contains(methods, Methods[2]),
		IsDelete:   slices.Contains(methods, Methods[3]),
		IsList:     slices.Contains(methods, Methods[4]),
		IsRetrieve: slices.Contains(methods, Methods[5]),
	}

	for _, l := range layers {
		if l == Layers[0] || l == Layers[4] {
			continue
		}
		layerPath, ok := LayerMap[l]
		if !ok {
			layerPath = l
		}
		filePath := path.Join(baseDir, "internal", layerPath, strcase.ToKebab(name)+".go")
		file, err := os.Create(filePath)
		if err != nil {
			if os.IsExist(err) {
				continue
			}
			panic(err)
		}
		p := path.Join(pkgDir, "..", "templates", l, "default.tpl") // root path
		paths := []string{p}
		if l != Layers[3] {
			for _, m := range methods {
				if m == Methods[0] {
					continue
				}
				p := path.Join(pkgDir, "..", "templates", l, m+".tpl")
				paths = append(paths, p)
			}
		}
		tmpl := template.
			Must(template.
				ParseFiles(paths...))
		if err := tmpl.ExecuteTemplate(file, "default", templateModel); err != nil {
			panic(err)
		}
		if l != Layers[3] {
			for _, m := range methods {
				if m == Methods[0] {
					continue
				}
				if err := tmpl.ExecuteTemplate(file, m, templateModel); err != nil {
					panic(err)
				}
			}
		}
	}
	if slices.Contains(layers, Layers[4]) {

		stepMap := make(map[string]bool)
		suite := &features.Suite{}
		definedSteps := suite.GetSteps()
		for key := range definedSteps {
			stepMap[key] = true
		}
		for _, m := range methods {
			if m == Methods[0] {
				continue
			}

			// generate features file
			p := path.Join(pkgDir, "..", "templates", Layers[4], m+".tpl")
			featureFilePath := path.Join(baseDir, "features", fmt.Sprintf("%s_%s.feature", m, strcase.ToKebab(name)))
			file, err := os.Create(featureFilePath)
			if err != nil {
				if os.IsExist(err) {
					continue
				}
				panic(err)
			}
			tmpl := template.
				Must(template.
					ParseFiles(p))
			if err := tmpl.ExecuteTemplate(file, m, templateModel); err != nil {
				panic(err)
			}

			// generate go file
			p = path.Join(pkgDir, "..", "templates", "godog", m+".tpl")
			filePath := path.Join(baseDir, "features", fmt.Sprintf("%s_%s.go", m, strcase.ToKebab(name)))
			file, err = os.Create(filePath)
			if err != nil {
				if os.IsExist(err) {
					continue
				}
				panic(err)
			}
			tmpl = template.
				Must(template.
					ParseFiles(p))
			if err := tmpl.ExecuteTemplate(file, m, templateModel); err != nil {
				panic(err)
			}

			steps := utils.GetStepsContent(featureFilePath, stepMap)
			bddFilePath := path.Join(baseDir, "features", "bdd.go")
			b, err := os.ReadFile(bddFilePath)
			if err != nil {
				panic(err)
			}
			steps = append(steps, "\n/*generate_key*/")
			replacer := strings.ReplaceAll(string(b), "/*generate_key*/", strings.Join(steps, ""))
			if err := ioutil.WriteFile(bddFilePath, []byte(replacer), fs.ModePerm); err != nil {
				panic(err)
			}
		}
	}
}

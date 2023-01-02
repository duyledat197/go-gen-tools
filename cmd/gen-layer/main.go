package main

import (
	"github.com/duyledat197/interview-hao/cmd/gen-layer/internal"
)

func main() {
	internal.Run()
}

// package main

// import (
// 	"io"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"text/template"
// )

// // templateFile defines the contents of a template to be stored in a file, for testing.
// type templateFile struct {
// 	name     string
// 	contents string
// }

// func createTestDir(files []templateFile) string {
// 	baseDir, _ := os.Getwd()

// 	dir, err := os.MkdirTemp(baseDir, "template")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, file := range files {
// 		f, err := os.Create(filepath.Join(dir, file.name))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer f.Close()
// 		_, err = io.WriteString(f, file.contents)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	return dir
// }

// func main() {
// 	// Here we create a temporary directory and populate it with our sample
// 	// template definition files; usually the template files would already
// 	// exist in some location known to the program.
// 	dir := createTestDir([]templateFile{
// 		// T0.tmpl is a plain template file that just invokes T1.
// 		{"T0.tmpl", `T0 invokes T1: ({{template "T1"}}) {{.Name}}`},
// 		// T1.tmpl defines a template, T1 that invokes T2.
// 		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
// 		// T2.tmpl defines a template T2.
// 		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
// 	})
// 	// Clean up after the test; another quirk of running as an example.
// 	// defer os.RemoveAll(dir)

// 	// pattern is the glob pattern used to find all the template files.
// 	pattern := filepath.Join(dir, "*.tmpl")
// 	type Name struct {
// 		Name string
// 	}
// 	// Here starts the example proper.
// 	// T0.tmpl is the first name matched, so it becomes the starting template,
// 	// the value returned by ParseGlob.
// 	tmpl := template.Must(template.ParseGlob(pattern))

// 	err := tmpl.Execute(os.Stdout, &Name{
// 		Name: "hello",
// 	})
// 	if err != nil {
// 		log.Fatalf("template execution: %s", err)
// 	}
// }

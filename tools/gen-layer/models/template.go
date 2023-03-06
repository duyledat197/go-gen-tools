package models

type Template struct {
	PascalCase string
	CamelCase  string
	Module     string
	IsCreate   bool
	IsRetrieve bool
	IsUpdate   bool
	IsList     bool
	IsDelete   bool
}

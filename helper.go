package main
import (
	"text/template"
	"strings"
)

var helper = template.FuncMap{
	"markdownTable": func(s string) (result string) {
		return strings.Replace(s, "\n", "<br>", -1)
	},
}
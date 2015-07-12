package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
	"text/template"
	"path/filepath"
)

const templateStringErrorTemplate = `
render template string failed

template:
%s

data:
%s

reason:
%s
`
const templateFileErrorTemplate = `
render template file

file:
%s

data:
%s

reason:
%s
`

func main() {
	app := cli.NewApp()
	app.Name = "gorender"
	app.Usage = "render go template on the fly"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "string, s",
			Usage: "template string",
		},
		cli.StringFlag{
			Name:  "path, p",
			Usage: "template path",
		},
	}

	app.Action = action
	app.Commands = []cli.Command{
		{
			Name: "csv",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "data-path",
				},
			},
			Action: csvRender,
		},
	}
	app.Run(os.Args)
}

func getTemplate(c *cli.Context) (t *template.Template) {
	templateString := c.GlobalString("string")
	templatePath := c.GlobalString("path")
	err := validateParameter(templateString, templatePath)
	if err != nil {
		log.Fatal(err)
	}

	if templateString != "" {
		return template.Must(template.New("").Funcs(helper).Parse(templateString))
	}

	if templatePath != "" {
		return template.Must(template.New(filepath.Base(templatePath)).Funcs(helper).ParseFiles(templatePath))
	}

	return
}

func render(t *template.Template, data map[string]string, out *os.File) (err error) {
	err = t.Execute(out, data)
	return
}

func validateParameter(templateString string, templatePath string) (err error) {
	if templatePath != "" && templateString != "" {
		return fmt.Errorf("do not use both string template and template file")
	}
	if templatePath == "" && templateString == "" {
		return fmt.Errorf("missing template")
	}
	return nil
}

func constructData(data *map[string]string, args cli.Args) (err error) {
	for _, x := range args {
		i := strings.Index(x, "=")
		if i == -1 {
			return fmt.Errorf("not a key value pair key=value: %s", x)
		}
		key := x[:i]
		value := x[i+1:]
		if key == "" {
			return fmt.Errorf("missing key for value %s", x)
		}
		(*data)[key] = value
	}
	return nil
}

func action(c *cli.Context) {
	t := getTemplate(c)

	data := map[string]string{}
	err := constructData(&data, c.Args())
	if err != nil {
		log.Fatalf("data is invalid, reason: %s", err)
	}

	render(t, data, os.Stdout)
}

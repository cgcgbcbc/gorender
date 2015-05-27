package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
	"text/template"
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
	app.Run(os.Args)
}

func renderStringTemplate(templateString string, data map[string]string) (err error) {
	tmpl, err := template.New("").Parse(templateString)
	if err != nil {
		return
	}
	return render(tmpl, data)
}

func renderFileTemplate(templatePath string, data map[string]string) (err error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return
	}
	return render(tmpl, data)
}

func render(t *template.Template, data map[string]string) (err error) {
	err = t.Execute(os.Stdout, data)
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
	templateString := c.String("string")
	templatePath := c.String("path")
	err := validateParameter(templateString, templatePath)
	if err != nil {
		log.Fatalf("parameter is invalid, reason: %s", err)
	}
	data := map[string]string{}
	err = constructData(&data, c.Args())
	if err != nil {
		log.Fatalf("data is invalid, reason: %s", err)
	}
	if templateString != "" {
		err := renderStringTemplate(templateString, data)
		if err != nil {
			log.Fatalf(templateStringErrorTemplate, templateString, data, err)
		}
	}
	if templatePath != "" {
		err := renderFileTemplate(templatePath, data)
		if err != nil {
			log.Fatalf(templateFileErrorTemplate, templatePath, data, err)
		}
	}
}

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
	"text/template"
)

func main() {
	app := cli.NewApp()
	app.Name = "gorender"
	app.Usage = "render go template on the fly"

	var templateString string
	var templatePath string

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

	app.Action = func(c *cli.Context) {
		templateString = c.String("string")
		templatePath = c.String("path")
		if templatePath != "" && templateString != "" {
			fmt.Println("do not use both string template and template file")
			return
		}
		if templatePath == "" && templateString == "" {
			fmt.Println("missing template")
			return
		}
		data := map[string]string{}
		for _, x := range c.Args() {
			i := strings.Index(x, "=")
			if i == -1 {
				fmt.Println("not a key value pair key=value,", x)
				return
			}
			key := x[:i]
			value := x[i+1:]
			if key == "" {
				fmt.Println("missing key,", x)
				return
			}
			data[key] = value
		}
		if templateString != "" {
			err := renderStringTemplate(templateString, data)
			if err != nil {
				panic(err)
			}
		}
		if templatePath != "" {
			err := renderFileTemplate(templatePath, data)
			if err != nil {
				panic(err)
			}
		}
	}

	app.Run(os.Args)
}

func renderStringTemplate(templateString string, data map[string]string) (err error) {
	tmpl, err := template.New("string").Parse(templateString)
	if err != nil {
		return
	}
	err = tmpl.Execute(os.Stdout, data)
	return
}

func renderFileTemplate(templatePath string, data map[string]string) (err error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return
	}
	err = tmpl.Execute(os.Stdout, data)
	return
}
